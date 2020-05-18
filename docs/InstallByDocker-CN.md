## 修改配置文件  
将文件[deployments/docker-compose/docker-compose.yml](deployments/docker-compose/docker-compose.yml)中的"localhostip"替换为系统物理网卡的真实ip。  

## 通过Docker-Compose启动Doraemon 
```bash
$ cd deployments/docker-compose/
$ docker-compose up -d
```
