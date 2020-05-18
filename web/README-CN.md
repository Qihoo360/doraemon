# 前端

### 安装好依赖包
```
// node -v >= 10.16.0
npm i
npm run dll
```
### 编译，热更新为开发模式
```
npm run dev
```

#### 涉及到的 api 域名指向需要修改 **app/config.js**
```
window.CONFIG = {
  // 这里修改为你的后端地址
  baseURL: 'http://localhost/',
}
```

### 打包命令

```
npm run build
```

### lint

```
npm run lint
```
