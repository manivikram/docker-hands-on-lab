# Distroless Health Application

## What the application does

This project runs a Go HTTP application that listens on port 8080.

The application exposes the following endpoint:

GET /health

Expected response:

```json
{"status":"UP"}
