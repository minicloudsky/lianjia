# lianjia

##### Translate to: [简体中文](README_zh.md)
## About Lianjia

[lianjia's](https://lianjia.com/) real estate data crawling and data analysis.

## Quickstart

### Deploy on Kubernetes
1. edit your configuration file config.yaml, config your mysql、redis、kafka... in ./deploy dir
2. create configmap
```shell
kubectl apply -f configmap.yaml
```
3. create deployment
```shell
kubectl apply -f deployment.yaml
```
4. create service
```shell
kubectl apply -f service.yaml
```

### Deploy on docker
1. edit your configuration file config.yaml, config your mysql、redis、kafka... in ./deploy dir
2. use docker compose to start your container
```shell
docker-compose up -d
```

## Develop

### Install Kratos
```
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```
### Create a service
```
# Create a template project
kratos new server

cd server
# Add a proto template
kratos proto add api/server/server.proto
# Generate the proto code
kratos proto client api/server/server.proto
# Generate the source code of service by proto file
kratos proto server api/server/server.proto -t internal/service

go generate ./...
go build -o ./bin/ ./...
./bin/server -conf ./configs
```
### Generate other auxiliary files by Makefile
```
# Download and update dependencies
make init
# Generate API files (include: pb.go, http, grpc, validate, swagger) by proto file
make api
# Generate all files
make all
```
### Automated Initialization (wire)
```
# install wire
go get github.com/google/wire/cmd/wire

# generate wire
cd cmd/server
wire
```

### Docker
```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8000:8000 -p 9000:9000 -v </path/to/your/configs>:/data/conf <your-docker-image-name>
```

