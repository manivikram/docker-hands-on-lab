Docker Image Comparison

Single-stage image
------------------
Image name: multistage-basic:single-stage
Image ID: 157c6755a2dd
Image size: 387 MB
Number of layers: 9
Build time: Add the actual build time from the time command

Multi-stage image
-----------------
Image name: multistage-basic:multi-stage
Image ID: e934141d4e0e
Image size: 15.2 MB
Number of layers: 3
Build time: Add the actual build time from the time command

Size comparison
---------------
The multi-stage image is 371.8 MB smaller than the single-stage image.

The multi-stage image is approximately 96% smaller.

Explanation
-----------
The single-stage image contains the Go compiler, build tools, source code,
dependencies, and the compiled application binary.

The multi-stage Dockerfile uses the Go image only in the build stage. The
compiled binary is copied into a lightweight Alpine runtime image.

The final multi-stage image does not contain the Go compiler, source code,
or unnecessary build tools. Therefore, it has fewer layers, uses less
storage, downloads faster, and has a smaller security attack surface.
