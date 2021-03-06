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


## the k8s way, locally

### start k8s

```shell
minikube start
```

### start an object

note: make sure you have your docker images built and pushed to docker hub

```shell
kubectl create -f cluster/app-pod.yml
```

```shell
kubectl create -f cluster/app-pod.yml --namespace simple-go
```

```shell
kubectl create -f cluster/app-namespace.yml
```

```shell
kubectl apply -f cluster
```

### forward a port

```shell
kubectl port-forward simple-go-pod 8080:8080
```

### check objects

```shell
kubectl get pods
```

```shell
kubectl get pods -w
```

```shell
kubectl get pods,deployments,ns
```

```shell
kubectl get pods --show-labels
```

```shell
kubectl get pods -o wide -L env
```

```shell
kubectl get pods --namespace simple-go
```

### add label

```shell
kubectl label pod simple-go-pod hello=world
```

```shell
kubectl label pod simple-go-pod env=prod --overwrite
```

### remove an object

```shell
kubectl delete -f cluster/app-pod.yml
```

```shell
kubectl delete pod simple-go-pod
```

```shell
kubectl delete -f cluster/app-namespace.yml
```

```shell
kubectl delete -f cluster
```

### scale

```shell
kubectl scale deployment simple-go --replicas=5
```

### check history

```shell
kubectl rollout history deployment simple-go
```

### update images

```shell
kubectl set image deployment simple-go simple-go=miloofcroton/simple-go:1.0.4
```

### undo image update

```shell
kubectl rollout undo deployment simple-go
```
