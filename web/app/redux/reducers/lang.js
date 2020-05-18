import { handleActions } from 'redux-actions'

export const langInfo = handleActions({
  'change lang'(state, action) {
    const { payload } = action
    return { ...state, lang: payload }
  },
}, { lang: '中文' })
