# 背景
源自 [360的开源项目doraemon](https://github.com/Qihoo360/doraemon)

但原项目源码维护落后于发布的镜像版本，且有些代码实现的略blabla，于是单独拉出来维护（welcome everybody）

## 原 README
[English](README.md) | [中文](README-CN.md)


# 对比 doraemon 和 alertmanager
## 结构对比

**doraemon = alertmanager + 一个 rule-engine**

doraemon 的动态配置告警规则能力，来自相对独立的 rule-engine：其负责实现 QueryFunc、NotifyFunc + 定期对 Promethues 进行reload

# 编译

```cassandraql
git clone https://github.com/huangwei2013/doraemon.git
cd doraemon
touch go.sum

make build-backend-image
make build-frontend-image
make build-ruleengine-image



```
# 额外说明
## 另：
  [这是一个相关项目](https://github.com/huangwei2013/myruleengine),延伸自 doraemon的 rule-engine，用于与 promethues-alertmanager 结合
  
  
## 再另：
  prometheus 是因为没有动态加载 rule 规则的能力，才有 360 这个项目的生存空间，所以。。。。。参看[这个](https://github.com/huangwei2013/prometheus)，基于prometheus的改造
