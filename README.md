# 背景
源自 [360的开源项目doraemon](https://github.com/Qihoo360/doraemon)

但原项目源码维护落后于发布的镜像版本，且有些代码实现的略blabla，于是单独拉出来维护（welcome everybody）

## 原 README
[English](README.md) | [中文](README-CN.md)


# 对比 doraemon 和 alertmanager
## 结构对比

**doraemon = alertmanager + 一个 rule-engine**

doraemon 的动态配置告警规则能力，来自相对独立的 rule-engine：其负责实现 QueryFunc、NotifyFunc + 定期对 Promethues 进行reload


-------------------------
另：
  [这是一个相关项目](https://github.com/huangwei2013/myruleengine),延伸自 doraemon的 rule-engine，用于与 promethues-alertmanager 结合

