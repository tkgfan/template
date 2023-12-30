FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV GOOS linux
ENV GOARCH amd64
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build
COPY . .
RUN go build -ldflags "-s -w" -o /app/app main.go

# 构建小镜像
FROM alpine:latest

# 设置时区
RUN apk add -U tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && apk del tzdata

WORKDIR /app
COPY --from=builder /app/app /app/app

EXPOSE 8888

CMD ["./app"]
