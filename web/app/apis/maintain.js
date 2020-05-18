
import { createApi } from '@ajax'

const prefix = '/api/v1';
/* maintains */
export const getMaintain = createApi(`${prefix}/maintains`, { method: 'get' }) // get list
export const addMaintain = createApi(`${prefix}/maintains`) // add
export const updateMaintain = createApi(`${prefix}/maintains/:id`, { method: 'put' }) // update
export const deleteMaintain = createApi(`${prefix}/maintains/:id`, { method: 'delete' }) // update
/* host */
export const getHost = createApi(`${prefix}/maintains/:id/hosts`, { method: 'get' }) // update
