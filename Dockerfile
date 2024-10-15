# 使用官方的 Golang 镜像作为构建环境
FROM golang:1.21-alpine as builder

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下载所有依赖
RUN go mod download

# 复制整个项目到容器中
COPY . .

# 构建可执行文件
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# 使用 alpine 镜像运行
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
