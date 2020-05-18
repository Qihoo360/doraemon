
import { createApi } from '@ajax'

const prefix = '/api/v1';
/* rules */
export const getRules = createApi(`${prefix}/rules`, { method: 'get' }) // get list
export const addRules = createApi(`${prefix}/rules`) // add 
export const updateRules = createApi(`${prefix}/rules/:id`, { method: 'put' }) // update
export const deleteRules = createApi(`${prefix}/rules/:id`, { method: 'delete' }) // update
