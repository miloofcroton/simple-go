# simple-go

## starting the server locally

### build

```shell
docker build -t miloofcroton/simple-go:1.0.1 .
```

or

```shell
GOOS=linux GOARCH=amd64 go build
docker build -t miloofcroton/simple-go:1.0.2 -f Dockerfile-v2 .
```

### push

```shell
docker push miloofcroton/simple-go:1.0.2
```

### run

(replace `9090` with any port number)

```shell
docker run \
    -it \
    -e "PORT=9090" \
    -p 9090:9090 \
    simple-go:1.0.1
```
