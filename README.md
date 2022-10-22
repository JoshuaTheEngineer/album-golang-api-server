# Album Golang API Server

Creating an API Server

## Libraries

1. [Golang](https://golang.google.cn/)
2. [Gin Web Framework](https://gin-gonic.com/)

## Features

### Healthcheck

To run a healthcheck

```
curl http://localhost:8080/health
```

### Get Album by ID

To get an album by ID

`{ID}` should be replaced with the album id

```
curl http://localhost:8080/albums/{ID} \
    --header "Content-Type: application/json" \
    --request "GET"
```

### List Albums

To list all albums

```
curl http://localhost:8080/albums \
    --header "Content-Type: application/json" \
    --request "GET"
```
