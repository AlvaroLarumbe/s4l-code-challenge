# StayForLong code challenge

## Requirements
- Golang 1.20 (probably it will work with 1.18+)
- Docker
- Docker Compose


## Project structure
This project follows the [Clean Architecture principles](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).  
The structure is as follows:

```
/cmd                App entrypoint
/application        Use cases
/domain             Domain objects
/infrastructure     HTTP server, HTTP handlers (controllers), request and response objects
/helper             Some helper tools out of pure Clean Architecture directories
/e2e                End to end tests
```


## Run
Start container using Docker Compose:

```sh
docker compose -p s4l-challenge up -d
```

This will create a Docker stack named `s4l-challenge` with a single container named `s4l-api`.  
This container exposes the API in the port **8090** of the host.

Alternative to run without Docker:

```sh
go run cmd/main.go
```


## Stop
Stop the stack using:

```sh
docker compose -p s4l-challenge stop
```


## Run tests
Unit and E2E tests:

```sh
# Assuming Unix OS and Go installed in host machine
go test ./...
```
