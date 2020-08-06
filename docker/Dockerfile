FROM golang:latest AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=off \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 将代码复制到容器中
ADD . $GOPATH/src/taskdash

# 移动到工作目录：/build
WORKDIR $GOPATH/src/taskdash

# 将我们的代码编译成二进制可执行文件 app

RUN go build -o taskdash $GOPATH/src/taskdash/app/main.go

EXPOSE 8080

# 需要运行的命令
ENTRYPOINT ["./taskdash"]
