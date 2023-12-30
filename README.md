# Robin

---

## Dev

### Prerequisites

- docker
- docker-compose

## Getting Started

You'll first need to copy `/dev/.env.template` to `/dev/.env` and fill out environment variables as needed.

From the root of the repo you can then run the following command to bring up all services:
```shell
docker compose -f dev/docker-compose.yml up
```

If this is the first time running the project you may need to build the images first:
```shell
docker compose -f dev/docker-compose.yml up --build
```

Once all services are up the API will be accessible from `http://127.0.0.1:3000/api`.
