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

### Add an Album

To add an album

```
curl http://localhost:8080/albums --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "{ID}", "title": "{TITLE}", "artist": "{ARTIST}", "price": {PRICE}}'
```

`{ID}` should be replaced with the album id

`{TITLE}` should be replaced with the album title

`{ARTIST}` should be replaced with the album artist

`{PRICE}` should be replaced with the album price in format _XX.XX_

### Delete an Album

To delete an album

```
curl http://localhost:8080/albums/{ID} \
    --header "Content-Type: application/json" \
    --request "DELETE"
```

`{ID}` should be replaced with the album id
