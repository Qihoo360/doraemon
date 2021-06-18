[English](README-EN.md) | [中文](README-CN.md)

# Doraemon

[![License](http://img.shields.io/badge/license-GPLv3+-brightgreen.svg)](LICENSE)

Doraemon is a **[Prometheus](https://prometheus.io)** based monitor system ,which are made up of three components——the Rule Engine,the Alert Gateway and the Web-UI.Instead of configuring alarm rules in config file,this system can configure alarm rules dynamically through the Web-UI and integrates many customized alarm functions.

## Features

- Users can configure alarm rules dynamically through the Web-UI.
- Support flexible alarm strategies such as alarm delays through which can realize the alarm upgrade strategies,alarm groups and duty groups.Users can handle the alarms in their own way by sending the alarms to hooks.
- Users can confirm the alarms by prometheus tags.
- Support the maintain groups.
- In order to reduce the number of alarms,all of which are aggregated by rules.The alarms are aggregated once per cycle and the alarm recovery information are aggregated every minute.
- LDAP/OAuth 2.0/DB Multiple login mode support for Enterprise Edition.

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
- Node.js v10.16+ and npm 6.10+ ([installation with nvm](https://github.com/creationix/nvm#usage))
- MySQL 5.6.X (Most of the data is in MySQL.)

## Quickly Start

- Clone

  ```bash
  $ git clone https://github.com/Qihoo360/doraemon.git
  ```

- Modify the Configuration File
  1.Replace the "localhost" in [deployments/docker-compose/conf/config.js](deployments/docker-compose/conf/config.js) with the local physical network card IP.
  2.Replace the "localhost" of WebUrl in [deployments/docker-compose/conf/app.conf](deployments/docker-compose/conf/app.conf) with the local physical network card IP.
- Start Doraemon

  Start server by docker-compose at Doraemon project.

  ```bash
  $ cd deployments/docker-compose/
  $ docker-compose up -d
  ```

  With the above command, you can access the Doraemon from http://hostip:32000. The default username is "admin",and the password is "123456".

## Instructions

**[Wiki](docs/readme.md)**

## Contributor

- [@BennieMeng](https://github.com/BennieMeng)
- [@JayRyu](https://github.com/jayryu)
- [@JoveYu](https://github.com/JoveYu)
- [@70data](https://github.com/70data)
