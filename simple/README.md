Simple
---

This is a simple RESTful API sample using Skaffold.

## Directory Structure
```
.
├── .dockerignore      # => files/directories to ignore from docker build and skaffold watch
├── Dockerfile
├── README.md
├── cmd
│    └── api
│        └── main.go
├── go.mod
├── handlers
│    ├── api.go
│    ├── greet.go
│    └── handler.go
├── k8s                # => k8s manifests
│    └── api.yaml
└── skaffold.yaml      # => Skaffold configs
```

## Usage

Run the application in minikube.

```shell
# start minikube
minikube start --profile simple

# treat any context as local
skaffold config set --global local-cluster true

# change the destination to the docker in minikube
eval $(minikube -p simple docker-env)

# run application
skaffold dev
```

Get url to access the service.

```shell
minikube service --profile simple greet-api --url
```

Stop the application.

```shell
# revert the destination to the local docker
eval $(minikube -p simple docker-env -u)

# stop minikube
minikube stop --profile simple

# delete minikube
minikube delete --profile simple
```
