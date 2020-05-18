# Rule Engine
[English](https://git.qihoo.cloud/sre/doraemon/blob/master/cmd/rule-engine/readme.md) | [中文](https://git.qihoo.cloud/sre/doraemon/blob/master/cmd/rule-engine/readme-CN.md)    
The core of rule engine is based on prometheus **[modules](https://github.com/prometheus/prometheus/rules)** , from which we separate the computing module and encapsulate it into a separate service.

## Functions

1. Request rule from the URL.Dynamically load rules and periodically reload them.
2. Read the data and calculate according the rules through the Prometheus's QueryAPI interface.
3. Send the alerts according to the NotifyURL.  

## Deployment  
### docker

```
docker run <image> --gateway.url=http://alert-gateway:port
```


