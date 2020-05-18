# 配置文件
### 简介
Doraemon默认的配置文件位于cmd/alert-gateway/conf/app.conf。

### 基本配置
*[参考Beego的配置](https://beego.me/docs/mvc/controller/config.md)*
- appname: App的名称
- httpport: http监听端口，默认是8080
- runmode: 运行模式，开发模式（dev）或者生产模式（prod）
- copyrequestbody: 是否允许在 HTTP 请求时，返回原始请求体数据字节，默认为 false （GET or HEAD or 上传文件请求除外） 
- autorender: 是否模板自动渲染，默认值为 true，对于 API 类型的应用，应用需要把该选项设置为 false，不需要渲染模板
- EnableDocs: 是否开启文档内置功能，默认是 false
- sessionon: 是否使用session，默认false

### 数据库配置
- DBTns: 数据库 Tns，示例：tcp(127.0.0.1:3306)
- DBName: 数据库名称 示例：doraemon
- DBUser: 数据库用户名，示例：root
- DBPasswd: 数据库密码，示例：root
- DBLoc: 数据库Location，示例：Asia%2FShanghai

### 相关接口配置
- SmsUrl: 短信发送接口，例如："http://127.0.0.1:20499/api/v1/sms"
- LanxinUrl: 蓝信发送接口，例如："http://127.0.0.1:20499/api/v1/lanxin/text"
- CallUrl: 电话接口，例如："http://127.0.0.1:20499/api/v1/lanxin/call"
- DutyGroupUrl: 获取值班组的接口，例如："http://127.0.0.1/Api/getDutyUser"
- BrokenUrl: 故障机器列表，例如："http://127.0.0.1/api/hosts/broken"

### Web-UI的域名
- WebUrl: Web-UI的域名，例如: "http://www.doraemon.com"