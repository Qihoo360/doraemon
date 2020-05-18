
const _toString = (value) => {
  if (typeof value === 'undefined' || Object.prototype.toString.call(value) === '[object Null]') {
    return false
  }
  if (typeof value === 'string') {
    return value
  } else if (typeof value !== 'object') {
    return value.toString()
  }
  return JSON.stringify(value)
}

export const setCookie = (k, v, day = 100) => {
  const key = _toString(k)
  const value = _toString(v)
  const exp = new Date();
  exp.setTime(exp.getTime() + (day * 24 * 60 * 60 * 1000));
  document.cookie = `${key}=${value};expires=${exp.toUTCString()}`;
}

export const localSet = (k, v) => {
  const key = _toString(k)
  const value = _toString(v)
  if (!key || !value) {
    return
  }
  if (window.localStorage) {
    try {
      localStorage.setItem(key, value)
    } catch (e) {
      setCookie(key, value)
    }
  } else {
    setCookie(key, value)
  }
}

export const getCookie = (k) => {
  const re = document.cookie.match(new RegExp(`(^| )${k}=([^;]*)(;|$)`))
  if (re) {
    return re[2]
  }
  return null;
}

export const localGet = (k) => {
  const key = _toString(k)
  if (!key) {
    return null
  }
  if (window.localStorage) {
    try {
      return localStorage.getItem(key)
    } catch (e) {
      return getCookie(key)
    }
  } else {
    return getCookie(key)
  }
}
