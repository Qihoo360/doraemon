import axios from 'axios'
// import { hashHistory } from 'react-router'
import { timeout, baseURL } from '@config'
import { message } from 'antd'
import { pathToRegexp, compile } from 'path-to-regexp'
import { Router } from '@configs/router.config'
import { HashRouter } from 'react-router-dom'
import { GlobalClass } from '@configs/common'

const { CancelToken } = axios

// 防止连续出现多个用户登录超时的提示
let flag = true
export function logOut(text, url) {
  console.log(text, flag)
  if (flag) {
    const { history } = GlobalClass
    if (history) {
      // 之前有 path='/' 页面，没有完成 history 复制，现在已经进行了重定向
      const state = encodeURIComponent(`${window.location.origin}${history.location.pathname}${history.location.search}`)
      console.log(history.location.pathname, String.prototype.startsWith.call(history.location.pathname, '/login'))
      if (!String.prototype.startsWith.call(history.location.pathname, '/login')) {
        history.push({
          pathname: '/login',
          search: `state=${state}`,
        })
        // window.location.href = `${Router === HashRouter ? '/#login' : '/login'}?state=${state}`
      }
    }
    flag = false
    setTimeout(() => flag = true, 0)
  }
}

let baseConfig = {
  url: '/',
  method: 'post', // default
  baseURL: '',
  headers: {
    'Content-Type': 'application/x-www-form-urlencoded',
  },
  params: {
    // ID: 12345,
  },
  data: {
  },
  timeout: '',
  withCredentials: true, // default
  responseType: 'json', // default
  maxContentLength: 2000,
  validateStatus(status) {
    return status >= 200 && status < 300 // default
  },
}

baseConfig = { ...baseConfig, timeout: timeout, baseURL: baseURL }

export const oftenFetchByPost = (api, options) => {
  // 当api参数为createApi创建的返回值
  if (typeof api === 'function') return api
  /**
   * 可用参数组合：
   * (data:Object,path:Object,sucess:Function,failure:Function,config:Object)
   * (data:Object,path:Object,sucess:Function,config:Object)
   * (data:Object,path:Object,sucess:Function)
   * (data:Object,path:Object,config:Object)
   * (data:Object)
   * ()
   */
  return (...rest) => { // 参数:(data:Object,path:Object,sucess?:Function,failure?:Function,config?:Object)
    // 参数分析
    let url = api
    const data = rest[0] || {}
    let start = 1
    const keys = []
    pathToRegexp(url, keys);
    if (keys.length) {
      const toPath = compile(url, { encode: encodeURIComponent })
      url = toPath({ ...rest[1] })
      start += 1
    }
    const token = sessionStorage.getItem('token')
    if (token) {
      // data.token = token
    }
    let success = null
    let failure = null
    let config = null
    for (let i = start; i < rest.length; i += 1) {
      if (typeof rest[i] === 'function') {
        if (!success) {
          success = rest[i]
        } else {
          failure = rest[i]
        }
      }
      if (Object.prototype.toString.call(rest[i]) === '[object Object]') {
        config = rest[i]
      }
    }

    const hooks = {
      abort: null,
    }

    const cancelToken = new CancelToken((c) => { hooks.abort = c })
    // 如果是用的30上的mock的服务，那么就默认不带cookie到服务器
    if (options && options.baseURL && (options.baseURL.indexOf('12602') !== -1)) {
      baseConfig.withCredentials = false
    } else {
      // 跨域呆cookie
      baseConfig.withCredentials = true
    }
    axios({
      ...baseConfig, ...options, ...config, url, data, cancelToken, headers: { ...baseConfig.headers, 'Content-Time': `${new Date().getTime()}` },
    })
      .then(response => response.data)
      .then((response) => {
        if (response === undefined || response === null) {
          message.error('没有返回数据')
          return
        }
        switch (response.code) {
          case 0: { success && success(response.data); break }
          case 1: {
            // message.warning(response.msg)
            // failure && failure(response)
            if (typeof failure === 'function') {
              failure(response)
            }
            message.error(response.msg)
            break
          }
          case -1: {
            logOut(response.msg)
            failure(response.msg)
            break
          }
          default: {
            if (typeof failure === 'function') {
              failure(response)
            } else {
              // logOut()
            }
          }
        }
      })
      .catch((e) => {
        if (axios.isCancel(e)) {
          if (process.env.NODE_ENV !== 'production') {
            console.log('Request canceled', e.message)
          }
        } else if (typeof failure === 'function') {
          if (e.code === 'ECONNABORTED') { // 超时的报错
            failure({
              data: '',
              msg: '服务器连接超时',
              code: -1,
            })
          } else {
            failure({
              data: '',
              msg: e.message,
              code: -1,
            })
          }
        }
      })
    return hooks
  }
}

// 创建发起api的启动器
export const createApi = function (api, options = {}) {
  // resolve api
  return oftenFetchByPost(`${api}`, options)
}

