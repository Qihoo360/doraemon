import { createApi } from '@ajax'

const prefix = '/api/v1';
/* groups */
export const getUser = createApi(`${prefix}/users`, { method: 'get' }) // get list
export const addUser = createApi(`${prefix}/users`) // add 
export const updatePassword = createApi(`${prefix}/users`, { method: 'put' }) // update
export const deleteUser = createApi(`${prefix}/users/:id`, { method: 'delete' }) // update