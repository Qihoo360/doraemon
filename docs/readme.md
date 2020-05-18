[English](https://github.com/Qihoo360/doraemon/blob/master/docs/readme.md) | [中文](https://github.com/Qihoo360/doraemon/blob/master/docs/readme-CN.md)  

# Introduction  

Doraemon is a **[Prometheus](https://prometheus.io)** based monitor system ,which are made up of three components——the Rule Engine,the Alert Gateway and the Web-UI.Instead of configuring alarm rules in config file,this system can configure alarm rules dynamically through the Web-UI and integrates many customized alarm functions. 

# Architecture  

![Architecture](images/Architecture.png)  

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

- **[User Installation Steps](https://github.com/Qihoo360/doraemon/blob/master/docs/UserInstallationSteps.md)**  
- **[Developer Installation Steps](https://github.com/Qihoo360/doraemon/blob/master/docs/DeveloperInstallationSteps.md)**

### Configuration File

- **[Configuration Item Description](https://github.com/Qihoo360/doraemon/blob/master/docs/ConfigurationItemDescription.md)**  

### Local User Managment

- **[The Default System User](https://github.com/Qihoo360/doraemon/blob/master/docs/DefaultUser.md)**
- **[Add Local User](https://github.com/Qihoo360/doraemon/blob/master/docs/AddUser.md)**
- **[Delete Local User](https://github.com/Qihoo360/doraemon/blob/master/docs/DeleteUser.md)**
- **[Change Password](https://github.com/Qihoo360/doraemon/blob/master/docs/ChangePassword.md)**

### Alarm System Instructions

- **[Create Alarm Plans and Strategies](https://github.com/Qihoo360/doraemon/blob/master/docs/CreateAlarmStrategies.md)**    
- **[Add Data Source](https://github.com/Qihoo360/doraemon/blob/master/docs/AddDataSource.md)**  
- **[Add Rules](https://github.com/Qihoo360/doraemon/blob/master/docs/AddRules.md)**  
- **[Add Alarm Group](https://github.com/Qihoo360/doraemon/blob/master/docs/AddAlarmGroup.md)**  
- **[Add Maintain Group](https://github.com/Qihoo360/doraemon/blob/master/docs/AddMaintainGroup.md)**  
- **[Confirm Alarms](https://github.com/Qihoo360/doraemon/blob/master/docs/ConfirmAlarms.md)**  
- **[View Historical Alarms](https://github.com/Qihoo360/doraemon/blob/master/docs/ViewHistoricalAlarms.md)**  
