
import { createApi } from '@ajax'

const prefix = '/api/v1';

export const getUserName = createApi(`${prefix}/login/username`, { method: 'get' }) // get list

export const getMethod = createApi(`${prefix}/login/method`, { method: 'get' })

export const localLoginApi = createApi(`${prefix}/login/local`)

export const ldapLoginApi = createApi(`${prefix}/login/ldap`)

export const oauthLoginApi = createApi(`${prefix}/login/oauth`, { method: 'get' })

export const logoff = createApi(`${prefix}/logout`, { method: 'get' })
