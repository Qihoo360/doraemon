
export const set = 'set$'
export const brandName = 'React' // slogan

// 开发环境默认配置
// let _serverIp = 'http://10.202.3.130'
// let _port = '16028'
// let _baseURL = `${_serverIp}:${_port}`
let _mockURL = 'http://localhost:1111/'
const _serverIp = 'http://localhost'
let _port = '8080'
// dev
// let _baseURL = 'http://10.213.116.29:8080/'
let _baseURL = window.CONFIG.baseURL
if (process.env.NODE_ENV === 'testing') { // 测试环境
  _mockURL = 'http://localhost:1111/'
  _port = '1111'
  _baseURL = `${_serverIp}:${_port}`
}
// if (process.env.NODE_ENV === 'production') { // 发布环境
//   _port = '16028'
//   _serverIp = 'http://10.202.3.130'
//   _baseURL = 'http://doraemon.qihoo.cloud/'
// }

export const serverIp = _serverIp
export const path = '/mock'
export const timeout = '15000' // 接口超时限制(ms)
export const baseURL = _baseURL
export const mockURL = _mockURL
