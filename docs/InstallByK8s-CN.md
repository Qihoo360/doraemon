# 使用Kubernetes部署  

## 从Github下载代码  
```shell
$ git clone https://github.com/Qihoo360/doraemon.git
```
## 安装依赖环境  
Doraemon依赖MySQL作为数据存储。  
```shell
$ kubectl apply -f deployments/kubernetes/mysql.yml
```
> 注意：这里使用容器启动MySQL，没有做数据持久化，在生产环境需要做数据持久化，以免数据丢失。

## 配置Configmap
1.为了配置系统的相关信息（比如数据库连接等等），需要根据[配置说明](ConfigurationItemDescription-CN.md)修改[deployments/kubernetes/doraemon.yml](../deployments/kubernetes/doraemon.yml)中的configmap。
> 如果使用[deployments/kubernetes/mysql.yml](../deployments/kubernetes/mysql.yml)中的配置来启动MySQL，就不需要改变configmap中的配置，系统会通过内部域名来连接MySQL。  

2.修改[deployments/kubernetes/doraemon.yml](../deployments/kubernetes/doraemon.yml)中doraemon-ui这个configmap，将"nodeip"替换为Kubernetes集群中任意节点的主机ip。

## 启动Doraemon
启动Doraemon前，请修改[deployments/kubernetes/doraemon.yml](../deployments/kubernetes/doraemon.yml)中的configmap，将"doraemon-ui"中的"baseURL"的"nodeip"替换为kubernetes集群中任意节点的节点ip。
```bash
$ kubectl apply -f deployments/kubernetes/doraemon.yml
```
现在用户可以通过 **http://nodeip:32000** 来访问Doraemon。  
