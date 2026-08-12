# Docker Named Volume Assignment

## What the application does

This project runs PostgreSQL using a Docker named volume.

A database and table are created with three records. The first container is
deleted, and a replacement container is started using the same volume.

The original records remain available, demonstrating persistent storage.

## Volume used

`postgres-data-volume`

## How to run

```bash
docker run -d \
  --name postgres-primary \
  -e POSTGRES_USER=devopsuser \
  -e POSTGRES_PASSWORD=DevOpsPass123 \
  -e POSTGRES_DB=companydb \
  -p 5433:5432 \
  -v postgres-data-volume:/var/lib/postgresql/data \
  postgres:16-alpine
