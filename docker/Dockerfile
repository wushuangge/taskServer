FROM golang:latest AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.io

WORKDIR /go/cache

ADD go.mod .
ADD go.sum .
RUN go mod download

# 移动到工作目录：/build
WORKDIR /go/src/taskdash

# 将代码复制到容器中
ADD . /go/src/taskdash

# 将我们的代码编译成二进制可执行文件 app
RUN go build -o taskdash /go/src/taskdash/app/main.go

# 需要运行的命令
ENTRYPOINT ["./taskdash"]