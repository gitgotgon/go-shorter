# go-shorter

A scalable URL shortener featuring short-link generation, redirection, analytics, caching, rate limiting, background processing and observabillity.

## Tech Stack

- Go
- PostgreSQL
- Goose
- Docker
- etc...

## Project Structure

- `cmd/api` - API application entry point
- `internal/config` - application configuration
- `internal/http` - HTTP routing and API components
- `migrations` - database schema migrations
- `deployments` - deployment and container configuration

## Setup

### Prerequisites

- Docker, Docker Compose

### Steps

1. copy `.env.example` to `.env`, change `POSTGRES_PASSWORD` to the password of your choice.
