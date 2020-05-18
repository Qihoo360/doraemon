import { createApi } from '@ajax'

const prefix = '/api/v1';
/* alerts */
export const getAlerts = createApi(`${prefix}/alerts`, { method: 'get' }) // get list

/* confirm */

export const getAlertsRules = createApi(`${prefix}/alerts/rules/:id`, { method: 'get' }) // get rules by id
export const confirmRules = createApi(`${prefix}/alerts`, { method: 'put' }) // confirm rules
export const getAllAlerts = createApi(`${prefix}/alerts/classify`, { method: 'get' }) // get all rules
