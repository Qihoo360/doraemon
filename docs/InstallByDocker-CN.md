## 克隆代码仓库

```bash
$ git clone https://github.com/Qihoo360/doraemon.git
```
## 修改配置文件  
1.将[deployments/docker-compose/conf/config.js](deployments/docker-compose/conf/config.js)中的"localhost"替换为本机物理网卡ip，端口号保持不变。  
2.修改[deployments/docker-compose/conf/app.conf](deployments/docker-compose/conf/app.conf)，将WebUrl中的"localhost"替换为本机物理网卡ip，端口号保持不变。其他配置选项见 **[配置项说明](ConfigurationItemDescription-CN.md)**。
## 启动服务  
在Doraemon的根目录下，通过 docker-compose 创建服务

```bash
$ cd deployments/docker-compose/
$ docker-compose up -d
```  
通过上述命令，您可以从通过 http://本机ip:32000 访问Doraemon。
