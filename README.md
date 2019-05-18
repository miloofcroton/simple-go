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

or

```shell
docker-compose build
```

### push

(not necessary with docker-compose)

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

or

```shell
docker-compose up
```


## the k8s way

### start k8s

```
minikube start
```

