FROM golang:1.21 AS builder

LABEL maintainer="minicloudsky <minicloudsky@gmail.com>"

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.cn GOOS=linux GOARCH=amd64 CGO_ENABLED=0 make build

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/bin /app
ENV TIME_ZONE=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone
RUN mkdir -p   /data/lianjia/logs/

WORKDIR /app

EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf

CMD ["./server", "-conf", "/data/conf"]
