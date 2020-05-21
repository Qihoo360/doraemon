## Clone

```bash
$ git clone https://github.com/Qihoo360/doraemon.git
```
## Modify the Configuration File  
1.Replace the "localhost" in [deployments/docker-compose/conf/config.js](../deployments/docker-compose/conf/config.js) with the local physical network card IP.  
2.Replace the "localhost" of WebUrl in [deployments/docker-compose/conf/app.conf](deployments/docker-compose/conf/app.conf) with the local physical network card IP.Description of other configurations can be found at **[Configuration Item Description](ConfigurationItemDescription.md)**.  
## Start Doraemon

  Start server by docker-compose at Doraemon project.

```bash
$ cd deployments/docker-compose/
$ docker-compose up -d
```

With the above command, you can access the Doraemon from http://hostip:32000. 
