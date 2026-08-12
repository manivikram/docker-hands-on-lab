# Regular vs Distroless Image Comparison

## Regular image

- Image: `distroless-health:regular`
- Base image: `alpine:3.20`
- Image size: 19.3 MB
- Number of filesystem layers: 4
- Shell available: Yes, `/bin/sh`
- Package manager available: Yes, `apk`
- Configured user: `appuser`
- Runs as root: No

## Distroless image

- Image: `distroless-health:distroless`
- Base image: `gcr.io/distroless/static-debian12:nonroot`
- Image size: 13.2 MB
- Number of filesystem layers: 14
- Shell available: No
- Package manager available: No
- Configured user: `nonroot:nonroot`
- Runs as root: No

## Size comparison

The distroless image is 6.1 MB smaller than the regular Alpine image.

The distroless image is approximately 32% smaller.

## Layer comparison

The regular image has 4 filesystem layers.

The distroless image has 14 filesystem layers.

Although the distroless image has more layers, it is still smaller. Image
size depends on the total contents of the layers, not only the number of
layers. The distroless base image is divided into multiple small runtime
layers but excludes general-purpose tools and packages.

## Health endpoint

Both images successfully returned:

```json
{"status":"UP"}



