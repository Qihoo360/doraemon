ARG GO_VERSION=1.12

FROM golang:${GO_VERSION}-alpine AS builder

ARG GOPROXY=https://goproxy.cn
ARG USER=docker

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk update && \
    env && \
    apk add git make && \
    rm -rf /var/cache/apk/*

WORKDIR /go/src/github.com/huangwei2013/doraemon/cmd/rule-engine

COPY . .

RUN export GO111MODULE=on && \
    export GOPROXY=https://mirrors.aliyun.com/goproxy/ && \
    go build cmd/rule-engine/main.go

FROM alpine:latest

ENV TZ=Asia/Shanghai

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk update && \
    apk add ca-certificates tzdata && \
    rm -rf /var/cache/apk/*

WORKDIR /data

COPY --from=builder /go/src/github.com/huangwei2013/doraemon/cmd/rule-engine/main ruleengine

ENTRYPOINT ["./ruleengine"]
