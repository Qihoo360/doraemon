[English](readme.md) | [中文](readme-CN.md)  
# 简介  
Doraemon是一个基于 **[Prometheus](https://prometheus.io)** 的监控系统。该系统主要由三个组件组成——计算引擎（Rule Engine），报警网关（Alert Gateway）以及Web-UI。与Prometheus使用静态的配置文件来配置报警规则不同，Doraemon可以通过Web-UI动态的配置加载报警规则。此外，Doraemon还集成了许多自定义的报警功能。

# 架构  
![Architecture](images/Architecture.png)  

# 术语  
- 报警规则: 与Prometheus中的 **[报警规则](https://prometheus.io/prometheus/latest/configuration/alerting_rules/)** 概念相同。
- 数据源: Prometheus Server的URL，由Rule Engine将报警规则下发至该URL进行计算。
- 报警接收组: 由多个报警接收人组成的组。 
- 值班组: 和报警接收组类似，但是它是动态的从接口中获取组成员的列表。
- 报警延迟: 报警触发一段时间后才将报警发送给接收人。
- 报警周期: 报警发送的周期。
- 报警计划: 由多条报警策略组成的集合。
- 报警方式: 对于内部用户，可以通过蓝信、短信和电话的方式进行报警。非内部用户可以采用HOOK的方式将报警转发到自定义的Web Server进行处理。
- 报警策略: 一条报警策略包含报警延迟、报警周期、报警时间段、报警接收组、值班组以及报警方式等配置信息。
- 报警确认: 如果需要短时间的暂停报警，可以通过勾选相应报警并填写暂停时长来确认报警。
- 维护组: 如果希望屏蔽一些固定时间段内某些特定机器的报警，可以通过配置报警维护组策略来实现。

# 使用文档  
### 安装步骤
- **[使用Docker-Compose部署](InstallByDocker-CN.md)**  
- **[使用Kubernetes部署](InstallByK8s-CN.md)**  

### 配置文件
- **[配置项说明](ConfigurationItemDescription-CN.md)**  

### 本地用户管理
- **[默认系统用户](DefaultUser-CN.md)**
- **[添加用户](AddUser-CN.md)**
- **[删除用户](DeleteUser-CN.md)**
- **[修改密码](ChangePassword-CN.md)**  

### 系统使用说明
- **[创建报警计划以及报警策略](CreateAlarmStrategies-CN.md)**    
- **[添加数据源](AddDataSource-CN.md)**  
- **[添加报警规则](AddRules-CN.md)**  
- **[添加报警接收组](AddAlarmGroup-CN.md)**  
- **[添加维护组](AddMaintainGroup-CN.md)**  
- **[报警确认](ConfirmAlarms-CN.md)**  
- **[查看历史报警](ViewHistoricalAlarms-CN.md)**  