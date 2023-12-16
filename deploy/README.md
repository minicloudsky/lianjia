# deploy lianjia
##### Translate to: [简体中文](README_zh.md)
You can deploy `lianjia` on Kubernetes or on docker with the following steps.

## On Kubernetes
1. edit your configuration file config.yaml, config your mysql、redis、kafka...
2. create namespace
```shell
kubectl create namespace lianjia
```
3create configmap
```shell
kubectl apply -f configmap.yaml
```
4create deployment
```shell
kubectl apply -f deployment.yaml
```
5create service
```shell
kubectl apply -f service.yaml
```

## On docker
1. edit your configuration file config.yaml, config your mysql、redis、kafka...
2. use docker compose to start your container
```shell
docker-compose up -d
```

## Tips
you can install mysql、redis、kafka with [helm](https://helm.sh/) as belows.

1. create `lianjia` namespace
```shell
kubectl create namespace lianjia
```
1. mysql
```shell
helm install lianjia-mysql-release oci://registry-1.docker.io/bitnamicharts/mysql --namespace lianjia --set image.tag=8.0
```

2. redis
```shell
helm install lianjia-redis-release oci://registry-1.docker.io/bitnamicharts/redis --namespace lianjia --set image.tag=5.0
```

3. kafka
```shell
helm install lianjia-kafka-release oci://registry-1.docker.io/bitnamicharts/kafka --namespace lianjia --set listeners.client.protocol='PLAINTEXT' --set listeners.controller.protocol='PLAINTEXT'
```

