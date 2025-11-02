# iris API

Backend API for the iris mobile application.

This repository contains the Go codebase that powers the public HTTP API, including routing, controllers, services, data models, and integration with external providers (Firebase, MongoDB/Firestore, etc.).


## Prerequisites
- Go 1.24+
- Docker (optional, for containerized runs)
- Access to a Firebase project and service account credentials (if not using mocks)

## Quick start

1) Configure the application
- Either edit config.yaml (local defaults) or provide environment variables described in docs/configs.md.
- If using real Firebase and database, place your service account JSON in a secure location and point FIREBASE_CREDENTIALS_FILE_PATH to it.

2) Run locally
- Native: go run ./main.go
- Docker: see docs/how_to/start_docker_image_locally.md

3) Health check
- Once started, the API listens on the configured port (default :8080). You can verify by calling a simple GET on any public endpoint you have configured.

## Development
- Lint: staticcheck ./...
- Tests: go test ./...
- Mocks: You can enable mocked Firebase/database via environment variables (see docs/configs.md) to run locally without external services.

## Observability
- Prometheus metrics are exposed on a dedicated port (see server.metrics_expose in docs/configs.md).

## Security
- Authentication is backed by Firebase JWT. For help generating a test token, see docs/how_to/generate_token_with_firebase.md.

## Versioning and releases
- Version is tracked in version.yaml and changes are documented in CHANGELOG.md.

## License
Proprietary — internal project of RoadTripMoustache / iris.
