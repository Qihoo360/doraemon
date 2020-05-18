
import { createApi } from '@ajax'

const prefix = '/api/v1';
/* groups */
export const getGroup = createApi(`${prefix}/groups`, { method: 'get' }) // get list
export const addGroup = createApi(`${prefix}/groups`) // add 
export const updateGroup = createApi(`${prefix}/groups/:id`, { method: 'put' }) // update
export const deleteGroup = createApi(`${prefix}/groups/:id`, { method: 'delete' }) // update
