FROM golang:1.12 as backend

WORKDIR /go/src/github.com/huangwei2013/doraemon/cmd/alert-gateway

COPY . .

RUN export GO111MODULE=on && \
    export GOPROXY=https://goproxy.cn && \
    go build cmd/alert-gateway/main.go && \
    mv main doraemon && \
    tar cvf pack.tar doraemon cmd/alert-gateway/conf/

FROM 360cloud/centos:7

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

COPY --from=backend /go/src/github.com/huangwei2013/doraemon/cmd/alert-gateway/pack.tar /opt/doraemon/

WORKDIR /opt/doraemon/

RUN tar -xvf pack.tar

CMD ["./doraemon"]
