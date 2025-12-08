# 构建阶段
FROM --platform=$BUILDPLATFORM golang:1.25-alpine AS builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

# 安装时区数据（用于后续复制）
RUN apk add --no-cache tzdata ca-certificates

# 设置工作目录
WORKDIR /app

# 设置 Go 代理（加速国内下载）
ENV GOPROXY=https://goproxy.cn,direct

# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用程序（交叉编译）
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -a -installsuffix cgo -ldflags="-s -w" -o main .

# 运行阶段 - 使用 distroless 镜像（更小更安全）
FROM gcr.io/distroless/static:nonroot

# 设置时区
ENV TZ=Asia/Shanghai

WORKDIR /app

# 从构建阶段复制时区数据
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# 从构建阶段复制 CA 证书
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# 从构建阶段复制编译好的二进制文件
COPY --from=builder /app/main .

# 暴露端口
EXPOSE 3061

# 使用非 root 用户运行
USER nonroot:nonroot

# 运行应用程序
ENTRYPOINT ["./main"]
