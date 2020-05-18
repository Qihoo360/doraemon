FROM golang:1.12 as backend

WORKDIR /go/src/github.com/Qihoo360/doraemon/cmd/alert-gateway

COPY cmd/alert-gateway .

COPY go.mod /go/src/github.com/Qihoo360/doraemon

COPY go.sum /go/src/github.com/Qihoo360/doraemon

RUN export GO111MODULE=on && \
    export GOPROXY=https://mirrors.aliyun.com/goproxy/ && \
    go build main.go && \
    mv main doraemon && \
    tar cvf pack.tar doraemon conf/

FROM 360cloud/centos:7

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

COPY --from=backend /go/src/github.com/Qihoo360/doraemon/cmd/alert-gateway/pack.tar /opt/doraemon/

WORKDIR /opt/doraemon/

RUN tar -xvf pack.tar

CMD ["./doraemon"]