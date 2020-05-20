[English](docs/readme.md) | [中文](docs/readme-CN.md)  
# Introduction  
Doraemon is a **[Prometheus](https://prometheus.io)** based monitor system ,which are made up of three components——the Rule Engine,the Alert Gateway and the Web-UI.Instead of configuring alarm rules in config file,this system can configure alarm rules dynamically through the Web-UI and integrates many customized alarm functions. 

# Architecture  
![Architecture](docs/images/Architecture.png)  

# Terminology  
- Rules:The same as **[alerting rules](https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/)** in prometheus.  
- Data Source:The endpoint of prometheus server to which the rules will be sent by the Rule Engine.
- Alarm Group:A set of username that the alarms will be sent to.
- Duty Group:The same as alarm group,however,it's taken from an interface dynamically. 
- Alarm Delays:A alarm will be sent to receivers after the time duration which is called alarm delays.
- Alarm period:The period of sending alarms to receivers.
- Alarm Plan:The set of alarm strategies.
- Alarm Method:For internal users,one can send alerts by LANXIN,SMS and CALL.The others can use the HOOK,which will send alarms to a http sever by which user can handle alarms in their own way.
- Alarm Strategy:Include the configuration of alarm delays,alarm period,alarm receivers,duty groups,alarm groups and alarm method.
- Confirm Alarms:If you want to stop receiving alarms in a short time,then you can confirm the alarms by filling in the duration.
- Maintain Group:If you want to stop receiving alarms triggered by some hosts within a certain period of time in a long time,you can create a maintain group for those hosts.


# Instructions  
### Installation Steps
- **[Deploy with Docker-Compose](docs/InstallByDocker.md)**  
- **[Deploy with Kubernetes](docs/InstallByK8s.md)**

### Configuration File
- **[Configuration Item Description](docs/ConfigurationItemDescription.md)**  

### Local User Managment
- **[The Default System User](docs/DefaultUser.md)**
- **[Add Local User](docs/AddUser.md)**
- **[Delete Local User](docs/DeleteUser.md)**
- **[Change Password](docs/ChangePassword.md)**

### Alarm System Instructions
- **[Create Alarm Plans and Strategies](docs/CreateAlarmStrategies.md)**    
- **[Add Data Source](docs/AddDataSource.md)**  
- **[Add Rules](docs/AddRules.md)**  
- **[Add Alarm Group](docs/AddAlarmGroup.md)**  
- **[Add Maintain Group](docs/AddMaintainGroup.md)**  
- **[Confirm Alarms](docs/ConfirmAlarms.md)**  
- **[View Historical Alarms](docs/ViewHistoricalAlarms.md)**  