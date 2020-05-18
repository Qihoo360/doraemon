import { routerReducer as routing } from 'react-router-redux'
import { combineReducers } from 'redux'

import * as tabList from './tabList'
import * as common from './common'
import * as demo from './demo'
import * as lang from './lang'

const rootReducer = combineReducers({
  routing,
  config: (state = {}) => state,
  ...tabList,
  ...common,
  ...demo,
  ...lang,
})

export default rootReducer
