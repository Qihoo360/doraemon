# Install with Kubernetes

## Clone Source Code from Github  
```shell
$ git clone https://github.com/Qihoo360/doraemon.git
```
## Install Dependent Environment   
Doraemon relies on MySQL,where MySQL is a required service.
```shell
$ kubectl apply -f deployments/kubernetes/mysql.yml
```
> Note:The data is not persisted.In production environment user should provide data persistence solution.

## Configure the Configmap    
1.To configurate related information such as database connection,user should modify the configmap in [deployments/kubernetes/doraemon.yml](deployments/kubernetes/doraemon.yml) according to the [instruction](ConfigurationItemDescription.md).
> If use [deployments/kubernetes/mysql.yml](deployments/kubernetes/mysql.yml) to startup MySQL,there is no need to modify the configmap,which connects MySQL through the inner domainname.  

2.Modify the configmap doraemon-ui in [deployments/kubernetes/doraemon.yml](deployments/kubernetes/doraemon.yml) by replacing the "nodeip" with host IP of any node in the ubernetes cluster.

## Startup Doraemon    
Before startup the Doraemon,you should modify the configmap in [deployments/kubernetes/doraemon.yml](deployments/kubernetes/doraemon.yml).Replace the "nodeip" of "baseURL" in "doraemon-ui" with the nodeip of any node in the kubernetes cluster.
```bash
$ kubectl apply -f deployments/kubernetes/doraemon.yml
```
Now user can visit Doraemon through **http://nodeip:32000**.  