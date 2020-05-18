import { createApi } from '@ajax'

const prefix = '/api/v1';

export const getPromethus = createApi(`${prefix}/proms`, { method: 'get' }) // get list
export const addPromethus = createApi(`${prefix}/proms`) // add
export const updatePromethus = createApi(`${prefix}/proms/:id`, { method: 'put' }) // update
export const deletePromethus = createApi(`${prefix}/proms/:id`, { method: 'delete' }) // delete
