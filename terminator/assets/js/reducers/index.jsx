import { combineReducers } from 'redux'
import releases from './releases'
import { routerReducer as routing } from 'react-router-redux'

const reducerFilter = combineReducers({
  releases,
  routing
})

export default reducerFilter
