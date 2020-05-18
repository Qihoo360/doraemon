# 使用kubernetes运行

## 从Github下载代码

```shell
$ git clone
```

## 安装依赖环境
Doraemon依赖MySQL作为数据存储。

```shell
$ kubectl apply -f deployments/kubernetes/mysql.yml
```

> 注意：这里使用容器启动MySQL，没有做数据持久化，在生产环境需要做数据持久化，以免数据丢失。

## 数据库初始化

Doraemon启动的时候，如果检测到数据库不存在，会自动创建数据库以及初始化数据。因此，用户不需要手动创建数据库。

## 配置Configmap

为了配置系统的相关信息（比如数据库连接等等），需要根据[配置说明](docs/ConfigurationItemDescription-CN.md)修改[deployments/kubernetes/doraemon.yml](deployments/kubernetes/doraemon.yml)中的configmap。
> 如果使用[deployments/kubernetes/mysql.yml](deployments/kubernetes/mysql.yml)中的配置来启动MySQL，就不需要改变configmap中的配置，系统会通过内部域名来连接MySQL。

## 启动Doraemon

启动Doraemon前，请修改[deployments/kubernetes/doraemon.yml](deployments/kubernetes/doraemon.yml)中的configmap，将"doraemon-ui"中的"baseURL"的"nodeip"替换为kubernetes集群中任意节点的节点ip。

```bash
$ kubectl apply -f deployments/kubernetes/doraemon.yml
```

现在用户可以通过 **http://nodeip:32000** 来访问Doraemon，系统默认的账户为admin:123456。  

## 启动RuleEngine  
```bash
$ kubectl apply -f deployments/kubernetes/rule-engine.yml
```
 
# 使用Docker-Compose安装

## 修改配置文件

将文件[deployments/docker-compose/conf/config.js](deployments/docker-compose/conf/config.js)中的"localhost"替换为当前主机理网卡的真实ip。  

## 通过Docker-Compose启动Doraemon 

```bash
$ cd deployments/docker-compose/
$ docker-compose up -d
```
