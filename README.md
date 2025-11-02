# Iris API

![Iris logo](./docs/logo.png)

![Docker Image Version](https://img.shields.io/docker/v/roadtripmoustache/iris-api?sort=semver)
[![GitHub stars](https://img.shields.io/github/stars/RoadTripMoustache/iris-api?style=social)](https://github.com/RoadTripMoustache/iris-api/stargazers)

**Iris API** provides a RESTful interface to create, retrieve, and manage ideas, votes, and comments from users.
It powers the feedback system of Iris, enabling efficient collection of bug reports and feature requests directly from your user base.

---

This repository contains the Go codebase that powers the public HTTP API, including routing, controllers, services, data models, and integration with external providers (Firebase, MongoDB/Firestore, etc.).


## Prerequisites
- Go 1.24+
- Docker (optional, for containerized runs)
- Access to a Firebase project and service account credentials (if not using mocks)

## Quick start

### Configure the application
Either edit config.yaml *(local defaults)* or provide environment variables described in [docs/configs.md](./docs/configs.md).

### Run locally
Use command `go run ./main.go`

### Health check
Once started, the API listens on the configured port *(default :8080)*. You can verify by calling a simple GET on any public endpoint you have configured.

To help you, you can either use the [openapi.yaml](./docs/openapi.yaml) contract or use directly the [Postman collection](./resources/Iris%20API.postman_collection.json)

## Observability
- Prometheus metrics are exposed on a dedicated port *(default :2121)* to help you monitore the API health.

---

## Contribution
If you want to contribute, you can do it by **opening an issue**, or [contribute in the project](./CONTRIBUTE.md).
