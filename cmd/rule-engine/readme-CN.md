[English](https://github.com/Qihoo360/doraemon/blob/master/cmd/rule-engine/readme.md) | [中文](https://github.com/Qihoo360/doraemon/blob/master/cmd/rule-engine/readme-CN.md)

# 计算引擎

核心使用 Prometheus 的 **[模块](https://github.com/prometheus/prometheus/rules)** 中的 Manager 来完成计算和告警，将计算模块单独剥离出来封装成一个独立的服务。

## 功能

1. 通过 URL 请求规则地址，动态加载规则并定期 reload。
2. 通过 Prometheus 的 QueryAPI 接口读取数据进行规则计算。
3. 按照配置的 NotifyURL 发送报警项。

## 部署运行

### docker

```
docker run <image> --gateway.url=http://alert-gateway:port
```
