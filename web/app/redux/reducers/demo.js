import { handleActions } from 'redux-actions'

export const demoResult = handleActions({
  'update demo'(state, action) {
    console.log(state, action)
    return { ...state, loading: true, test: true }
  },
}, { init: 0 })
