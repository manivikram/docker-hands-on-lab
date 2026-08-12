# docker-hands-on-lab

# Docker Hands-On Lab

## About This Repository

I created this repository to document my hands-on Docker practice.

Instead of only learning Docker commands individually, I worked through small assignments that helped me understand how Docker behaves in practical scenarios such as building images, running multiple application versions, reducing image size, persisting data, and sharing files between the host and containers.

I completed the exercises on an Ubuntu EC2 instance and recorded the Docker commands, observations, test results, and project documentation inside each assignment directory.

---

## What I Practiced

Through these assignments, I worked with:

* Docker images and containers
* Dockerfiles
* Multi-stage builds
* Image layers and image-size optimization
* React applications with Nginx
* Go applications
* Distroless container images
* Non-root containers
* Docker bind mounts
* Docker named volumes
* PostgreSQL containers
* Persistent database storage
* Port mapping
* Container inspection
* Docker networking concepts
* Container troubleshooting

---

# Repository Structure

```text
docker-hands-on-lab/
│
├── multistage-project/
│   ├── main.go
│   ├── go.mod
│   ├── Dockerfile
│   ├── Dockerfile.single-stage
│   ├── README.md
│   ├── commands.txt
│   └── comparison.txt
│
├── multistage-web/
│   ├── src/
│   ├── index.html
│   ├── package.json
│   ├── vite.config.js
│   ├── Dockerfile
│   ├── .dockerignore
│   ├── README.md
│   ├── commands.txt
│   └── test-results.txt
│
├── distroless-app/
│   ├── main.go
│   ├── go.mod
│   ├── Dockerfile
│   ├── Dockerfile.regular
│   ├── README.md
│   ├── commands.txt
│   └── image-comparison.md
│
├── bind-mount/
│   ├── website/
│   │   └── index.html
│   ├── README.md
│   ├── commands.txt
│   └── observations.txt
│
└── docker_named-volume/
    ├── README.md
    ├── commands.txt
    ├── database-commands.txt
    └── persistence-test.txt
```

---

# Project 1 - Multi-Stage Docker Build

## Objective

The goal of this assignment was to understand the difference between a normal single-stage Docker build and a multi-stage Docker build.

I created a small Go application that prints:

```text
Welcome to the Docker Multi-Stage Build Assignment
```

I then created two Docker images.

### Single-Stage Build

The single-stage image contained:

```text
Go compiler
+
Source code
+
Build tools
+
Compiled binary
```

### Multi-Stage Build

The multi-stage build separated compilation from runtime.

```text
              BUILD STAGE

             golang image
                  |
                  v
            Compile Go App
                  |
                  v
           Application Binary
                  |
                  |
                  v
             RUNTIME STAGE

               Alpine
                  |
                  v
           Application Binary
                  |
                  v
             Container
```

Only the compiled binary was copied into the runtime image.

## Results

```text
Single-stage image
Size:   387 MB
Layers: 9

Multi-stage image
Size:   15.2 MB
Layers: 3
```

The multi-stage image was approximately 96% smaller.

This assignment helped me understand why build tools do not need to remain inside production container images.

---

# Project 2 - Multi-Stage Web Application

## Objective

The goal of this assignment was to build a web application using one image for the build process and a lightweight image for production.

I created a React application and used:

```text
Node.js
   |
   | npm install
   | npm run build
   v
React Production Files
   |
   v
Nginx
   |
   v
Browser
```

The Node.js image was used only during the build stage.

The final runtime image contained Nginx and the generated website files.

---

## Running Two Versions

I created two versions of the application.

```text
Version 1.0
Container: multistage-web-v1
Host Port: 8081

Version 2.0
Container: multistage-web-v2
Host Port: 8082
```

Both versions were running at the same time.

```text
Browser
   |
   +------ :8081 ------> Version 1.0
   |
   +------ :8082 ------> Version 2.0
```

This helped me understand Docker image tags, container names, port mapping, and application versioning.

---

# Project 3 - Distroless Container

## Objective

The goal of this assignment was to understand how distroless images can reduce unnecessary software inside production containers.

I created a Go HTTP application exposing:

```text
GET /health
```

The expected response was:

```json
{"status":"UP"}
```

I built two versions:

```text
distroless-health:regular
distroless-health:distroless
```

---

## Regular Container

The regular container used Alpine.

It contained:

```text
Application
+
Shell
+
apk package manager
+
Linux utilities
```

I verified that I could enter the container with:

```bash
docker exec -it health-regular /bin/sh
```

I also confirmed that the package manager existed:

```bash
docker exec health-regular which apk
```

---

## Distroless Container

The distroless image contained only the minimum runtime components required by the application.

When I attempted:

```bash
docker exec -it health-distroless /bin/sh
```

Docker returned an error because `/bin/sh` did not exist.

The package-manager test also failed because `apk` was not available.

---

## Image Comparison

```text
Regular image
Size:   19.3 MB
Layers: 4

Distroless image
Size:   13.2 MB
Layers: 14
```

The distroless image had more filesystem layers but was still smaller.

This helped me understand that the number of layers alone does not determine image size.

I also configured both containers to run as non-root users.

---

# Project 4 - Docker Bind Mount

## Objective

The purpose of this assignment was to understand how a host directory can be shared directly with a running container.

I created:

```text
website/
└── index.html
```

and mounted it into an Nginx container.

```text
Ubuntu EC2 Host

website/index.html
        |
        | Bind Mount
        v
Nginx Container

/usr/share/nginx/html/index.html
        |
        v
Browser
```

The container was started with a bind mount similar to:

```bash
docker run -d \
  --name bind-mount-nginx \
  -p 8083:80 \
  -v $(pwd)/website:/usr/share/nginx/html \
  nginx
```

Initially the webpage displayed:

```text
Docker Bind Mount - Version 1
```

I then modified `index.html` directly on the EC2 host.

After refreshing the browser, it displayed:

```text
Docker Bind Mount - Version 2
```

I did not rebuild the Docker image or restart the container.

This demonstrated that bind mounts provide direct access to host files.

---

# Project 5 - Docker Named Volume

## Objective

The purpose of this assignment was to understand how Docker volumes preserve data independently of the container lifecycle.

I used PostgreSQL for this exercise.

The architecture was:

```text
PostgreSQL Container
        |
        v
postgres-data-volume
        |
        v
Persistent Database Data
```

I created:

```text
Database: companydb
Table:    employees
```

and inserted multiple records.

I then:

1. Stopped the original PostgreSQL container.
2. Removed the container.
3. Kept the named volume.
4. Created a new PostgreSQL container.
5. Attached the same volume.
6. Queried the database again.

The original table and records were still available.

---

## Persistence Flow

```text
postgres-primary
       |
       v
postgres-data-volume
       |
       | Container deleted
       |
       v
Volume still exists
       |
       v
postgres-replacement
       |
       v
Existing database available
```

This showed me that:

```text
Container lifecycle != Data lifecycle
```

Deleting a container does not delete data stored in a named volume.

---

## Docker Volume Location

I inspected the volume using:

```bash
docker volume inspect postgres-data-volume
```

Docker reported the volume location as:

```text
/var/lib/docker/volumes/postgres-data-volume/_data
```

I treated this as Docker-managed storage and did not manually modify the PostgreSQL files inside it.

---

# Commands I Used Frequently

Some of the Docker commands I practiced throughout these assignments include:

```bash
docker pull
docker build
docker run
docker ps
docker ps -a
docker stop
docker start
docker rm
docker images
docker image ls
docker history
docker image inspect
docker inspect
docker exec
docker top
docker logs
docker volume create
docker volume ls
docker volume inspect
docker system df
docker system prune
```

For testing web applications I also used:

```bash
curl
```

For checking ports on the EC2 instance I used commands such as:

```bash
ss -tulnp
netstat -tulnp
lsof -i -P -n
```

---

# Port Mapping

One concept I used repeatedly was Docker port mapping.

For example:

```bash
docker run -p 8081:80 nginx
```

means:

```text
Host / EC2
Port 8081
    |
    v
Docker Container
Port 80
```

From inside the EC2 instance I could test using:

```bash
curl http://localhost:8081
```

From my laptop browser I used:

```text
http://<EC2-PUBLIC-IP>:8081
```

This helped me understand the difference between the Docker host port and the container port.

---

# What I Learned

Working through these assignments helped me understand Docker as more than just `docker run`.

I practiced the complete workflow:

```text
Application
     |
     v
Dockerfile
     |
     v
docker build
     |
     v
Docker Image
     |
     v
docker run
     |
     v
Container
     |
     +---- Port Mapping
     |
     +---- Bind Mount
     |
     +---- Named Volume
     |
     +---- Networking
     |
     +---- Logs / Inspect
```

The most important concepts I learned were:

* Why multi-stage builds reduce final image size.
* Why build tools should not remain in production images.
* How Nginx can serve a frontend application built using Node.js.
* How image tags can be used to maintain multiple application versions.
* Why distroless images have a smaller attack surface.
* Why production containers should run as non-root users where possible.
* How bind mounts expose host files directly to containers.
* How named volumes preserve data after container deletion.
* How host ports map to container ports.
* How to inspect running containers and troubleshoot Docker issues.

---

# Final Result

These labs gave me hands-on experience with Docker image creation, container management, storage, application deployment, security, and troubleshooting.

Instead of only reading Docker concepts, I built and tested each scenario on an Ubuntu EC2 instance and documented the commands and results inside each project directory.
