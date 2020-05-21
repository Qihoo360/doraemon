[English](readme.md) | [中文](readme-CN.md)  
# Introduction  
Doraemon is a **[Prometheus](https://prometheus.io)** based monitor system ,which are made up of three components——the Rule Engine,the Alert Gateway and the Web-UI.Instead of configuring alarm rules in config file,this system can configure alarm rules dynamically through the Web-UI and integrates many customized alarm functions. 

# Architecture  
![Architecture](images/Architecture.png)  

# Terminology  
- Rules:The same as **[alerting rules](https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/#alerting-rules)** in prometheus.  
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
- **[Deploy with Docker-Compose](InstallByDocker.md)**  
- **[Deploy with Kubernetes](InstallByK8s.md)**

### Configuration File
- **[Configuration Item Description](ConfigurationItemDescription.md)**  

### Local User Managment
- **[The Default System User](DefaultUser.md)**
- **[Add Local User](AddUser.md)**
- **[Delete Local User](DeleteUser.md)**
- **[Change Password](ChangePassword.md)**

### Alarm System Instructions
- **[Create Alarm Plans and Strategies](CreateAlarmStrategies.md)**    
- **[Add Data Source](AddDataSource.md)**  
- **[Add Rules](AddRules.md)**  
- **[Add Alarm Group](AddAlarmGroup.md)**  
- **[Add Maintain Group](AddMaintainGroup.md)**  
- **[Confirm Alarms](ConfirmAlarms.md)**  
- **[View Historical Alarms](ViewHistoricalAlarms.md)**  
