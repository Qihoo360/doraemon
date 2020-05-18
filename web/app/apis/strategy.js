import { createApi } from '@ajax'

const prefix = '/api/v1';
/* strategy */
export const getStrategy = createApi(`${prefix}/plans`, { method: 'get' }) // get list
export const addStrategy = createApi(`${prefix}/plans`) // add 
export const updateStrategy = createApi(`${prefix}/plans/:id`, { method: 'put' }) // update
export const deleteStrategy = createApi(`${prefix}/plans/:id`, { method: 'delete' }) // update

/* receiver */

export const getReceiver = createApi(`${prefix}/plans/:id/receivers`, { method: 'get' }) // get list
export const addReceiver = createApi(`${prefix}/plans/:id/receivers`) // add 
export const updateReceiver = createApi(`${prefix}/receivers/:id`, { method: 'put' }) // update
export const deleteReceiver = createApi(`${prefix}/receivers/:id`, { method: 'delete' }) // update