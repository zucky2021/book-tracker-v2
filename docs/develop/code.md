# Code

## Manage version

[Github](https://github.com/zucky2021/book-tracker-v2)

## Docker memo

- After editing Dockerfile or docker-compose.yml, run this code.

```bash
docker compose down && docker compose up --build -d
```

- Cache disable

```bash
docker-compose down --rmi all --volumes --remove-orphans
docker-compose up --build -d
```
