# 基础镜像
FROM golang:latest as build_stage

# 容器环境配置
ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE="on"

#作者
MAINTAINER ukinhappy

#工作区
WORKDIR /go/src/app

#源代码拷贝到仓库中
COPY . .

# 编译
RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o echoserver


#部署
#部署到最小linux镜像中
FROM alpine:latest

WORKDIR /root/

#从编译阶段 复制文件到当前目录
COPY --from=build_stage /go/src/app/echoserver .

#端口报漏
EXPOSE 8081

# 启动服务
ENTRYPOINT ["./echoserver"]