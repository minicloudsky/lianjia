# 部署链家

你可以通过以下步骤在 Kubernetes 或 Docker 上部署 `lianjia`。

## 在 Kubernetes 上

1. 编辑你的配置文件 `config.yaml`，配置你的 MySQL、Redis、Kafka...
2. 创建命名空间：

```shell
kubectl create namespace lianjia
```

3. 创建 ConfigMap：

```shell
kubectl apply -f configmap.yaml
```

4. 创建 Deployment：

```shell
kubectl apply -f deployment.yaml
```

5. 创建服务：

```shell
kubectl apply -f service.yaml
```

## 在 Docker 上

1. 编辑你的配置文件 `config.yaml`，配置你的 MySQL、Redis、Kafka...
2. 使用 Docker Compose 启动容器：

```shell
docker-compose up -d
```

## 提示

你可以使用 [Helm](https://helm.sh/) 安装 MySQL、Redis、Kafka，具体步骤如下。

1. 创建 `lianjia` 命名空间：

```shell
kubectl create namespace lianjia
```

2. 安装 MySQL：

```shell
helm install lianjia-mysql-release oci://registry-1.docker.io/bitnamicharts/mysql --namespace lianjia --set image.tag=8.0
```

3. 安装 Redis：

```shell
helm install lianjia-redis-release oci://registry-1.docker.io/bitnamicharts/redis --namespace lianjia --set image.tag=5.0
```

4. 安装 Kafka：

```shell
helm install lianjia-kafka-release oci://registry-1.docker.io/bitnamicharts/kafka --namespace lianjia --set listeners.client.protocol='PLAINTEXT' --set listeners.controller.protocol='PLAINTEXT'
```