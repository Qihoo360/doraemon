
import { createAction } from 'redux-actions'
import * as login from '@apis/login'
import { getPromethus } from '@apis/promethus'
import { getStrategy } from '@apis/strategy'
import { createAjaxAction } from '@configs/common'


// login 登陆
export const requestLogin = createAction('request login')
export const recevieLogin = createAction('receive login')
export const rejectLogin = createAction('reject login')
export const loginAction = createAjaxAction(login.getUserName, requestLogin, recevieLogin, rejectLogin)

// promethus
export const requestPromethus = createAction('request promethus')
export const receivePromethus = createAction('receive promethus')
export const promethus = createAjaxAction(getPromethus, requestPromethus, receivePromethus)

// promethus
export const requestStrategy = createAction('request strategy')
export const receiveStrategy = createAction('receive strategy')
export const strategy = createAjaxAction(getStrategy, requestStrategy, receiveStrategy)

// gFormCache gfor2.0m的缓存
export const setGformCache2 = createAction('set gform cache2')
export const clearGformCache2 = createAction('clear gform cache2')

// socket receive
// export const socketReceive = createAction('socketReceive')
