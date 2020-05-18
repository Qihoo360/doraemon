[English](https://github.com/Qihoo360/doraemon/blob/master/README.md) | [中文](https://github.com/Qihoo360/doraemon/blob/master/README-CN.md)  

# Doraemon

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/Qihoo360/doraemon/blob/master/LICENSE)

Doraemon is a **[Prometheus](https://prometheus.io)** based monitor system ,which are made up of three components——the Rule Engine,the Alert Gateway and the Web-UI.Instead of configuring alarm rules in config file,this system can configure alarm rules dynamically through the Web-UI and integrates many customized alarm functions. 

## Features

- Users can configure alarm rules dynamically through the Web-UI.
- Support flexible alarm strategies such as alarm delays through which can realize the alarm upgrade strategies,alarm groups and duty groups.Users can handle the alarms in their own way by sending the alarms to hooks.
- Users can confirm the alarms by prometheus tags.
- Support the maintain groups.
- In order to reduce the number of alarms,all of which are aggregated by rules.The alarms are aggregated once per cycle and the alarm recovery information are aggregated every minute.
- LDAP/OAuth 2.0/DB Multiple login mode support.

## Architecture
The whole system adopts the separation of front and back ends, in which the front end uses React for data interaction and display.The backend uses the **[Beego](https://beego.me)** framework for data interface processing and data for MySQL storage.  
  
![Architecture](docs/images/Architecture.png)  

## Component
- Rule Engine:Pull rules from Alert Gateway,and then send the rules to prometheus server to caculate and push the alerts to Alert Gateway.
- Alert Gateway:Aggregate the alarms and send them to alarm receivers according to their alarm strategies.
- Web UI:For adding rules,alarm strategies and maintain groups.To confirm alarms and view historical alarm records.

## Dependence

- Golang 1.12+ ([installation manual](https://golang.org/dl/))
- Docker 17.05+ ([installation manual](https://docs.docker.com/install))
- Bee ([installation manual](https://github.com/beego/bee))
- Node.js v11+ and npm 6.5+ ([installation with nvm](https://github.com/creationix/nvm#usage))
- MySQL 5.6+ (Most of the data is in MySQL.)

## Quickly Start

- Clone

```bash
$ go get 
```

- Start Doraemon

  Start server by docker-compose at Doraemon project.

```bash
$ cd deployments/docker-compose/
$ docker-compose up -d
```

With the above command, you can access the local Doraemon from http://127.0.0.1:4200, the default administrator account admin:123456.  

## Instructions

**[Wiki](https://github.com/Qihoo360/doraemon/README)**

## Contributor  

- [@BennieMeng](https://github.com/BennieMeng)  
- [@JayRyu](https://github.com/jayryu)  
- [@JoveYu](https://github.com/JoveYu)
- [@70data](https://github.com/70data)
