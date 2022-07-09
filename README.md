# goat

goat is a go api template :wink:

## run with docker

```bash
docker build -t goat .
docker run -p 8080:8080 goat
```

## run with compose

```bash
docker-compose up -d
```
show logs:
```bash
docker-compose logs goat
docker-compose logs postgres
```

## run compose in dev mode

> in dev mode the air package is used for hot reload

```bash
docker-compose -f docker-compose.dev.yml up --build
```
