# lianjia

##### Translate to: [English](README.md)
[lianjia](https://lianjia.com/) 链家网房产数据爬取和数据分析.

## 快速开始

### Kubernetes 部署
1. 编辑你的 config.yaml, 配置在./deploy目录下你的 mysql、redis、kafka... 服务
2. 创建 k8s namespace
```shell
kubectl create namespace lianjia
```
3. 创建 configmap
```shell
kubectl apply -f configmap.yaml
```
4. 创建 deployment
```shell
kubectl apply -f deployment.yaml
```
5. 创建 service
```shell
kubectl apply -f service.yaml
```

### Docker 部署
1. 编辑你的 config.yaml, 配置在./deploy目录下你的 mysql、redis、kafka... 服务
2. 使用 `docker-compose` 启动服务
```shell
docker-compose up -d
```

## 开发

### 安装 Kratos
```
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```
### 创建服务
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
### 通过 Makefile 生成 protobuf 对应的代码
```
# Download and update dependencies
make init
# Generate API files (include: pb.go, http, grpc, validate, swagger) by proto file
make api
# Generate all files
make all
```
### 安装 wire
```
# install wire
go get github.com/google/wire/cmd/wire

# 生成依赖注入
cd cmd/server
wire
```

### docker 镜像构建
```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8000:8000 -p 9000:9000 -v </path/to/your/configs>:/data/conf <your-docker-image-name>
```

