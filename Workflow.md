
== Building the Go microservices

```shell
go get -d ./...
go build
./Simple-go.exe
```

== Containerization with Docker

```shell
docker build -t simple-go:1.0.1 .
docker images
docker tag simple-go:1.0.1 miloofcroton/simple-go:1.0.1
docker push miloofcroton/simple-go:1.0.1
```

== Running Docker image locally

```shell
docker run -it -p 8080:8080 simple-go:1.0.1
docker run -it -e "PORT=9090" -p 9090:9090 simple-go:1.0.1
docker ps --all

docker run --name=simple-go -d -p 8080:8080 simple-go:1.0.1
docker ps
docker stats
docker stop
docker start
```

== Improved Containerization with Docker

```shell
GOOS=linux GOARCH=amd64 go build
docker build -t simple-go:1.0.1-alpine .
docker images
docker tag simple-go:1.0.1-alpine miloofcroton/simple-go:1.0.1-alpine
docker push miloofcroton/simple-go:1.0.1-alpine
```

== Docker Compose

```shell
docker-compose build
docker images
docker-compose up -d
docker-compose kill
docker-compose rm
```

== Kubernetes and Pods

```shell
kubectl get pods
kubectl create -f k8s-pod.yml
kubectl get pods
kubectl describe pod simple-go

kubectl port-forward simple-go 8080:8080

kubectl get pods --show-labels
kubectl get pods -o wide -L env
kubectl label pod simple-go hello=world
kubectl get pods -o wide -L hello
kubectl label pod simple-go env=prod --overwrite
kubectl get pods -o wide -L env

kubectl get ns
kubectl get pods --namespace kube-system
kubectl create -f k8s-namespace.yml
kubectl get ns
kubectl create -f k8s-pod.yml --namespace simple-go

kubectl delete -f k8s-pod.yml
```

== Kubernetes Deployments and Services

```shell
kubectl get services,deployments,pods

kubectl create -f k8s-deployment.yml
kubectl get deployments,pods

kubectl apply -f k8s-deployment.yml

kubectl create -f k8s-service.yml
kubectl get services
kubectl describe service simple-go

kubectl delete -f k8s-deployment.yml
kubectl delete -f k8s-service.yml
```

== Kubernetes Scaling and Rolling Updates

```shell
kubectl create -f k8s-deployment.yml --record=true

kubectl scale deployment simple-go --replicas=5
kubectl scale deployment simple-go --replicas=3

kubectl rollout history deployment simple-go

kubectl apply -f k8s-deployment.yml

kubectl edit deployment simple-go

kubectl set image deployment simple-go simple-go=simple-go:1.0.2

kubectl rollout history deployment simple-go
kubectl rollout undo deployment simple-go --to-revision=0

kubectl delete -f k8s-deployment.yml
```
