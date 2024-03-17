# Test task: Auth service API

## Requirements

- Golang 1.22.0
- Docker and Docker Compose
- Make

## Prepare

Copy env variables from `.env.dist` to `.env` file:

```bash
cp .env.dist .env
```

Run MongoDB from Docker Compose file:

```bash
docker-compose up -d
```

Install dependencies:

```bash
make install
```

## Run

Start server:

```bash
make run
```
