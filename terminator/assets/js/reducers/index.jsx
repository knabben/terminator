import { combineReducers } from 'redux'
import visibilityFilter from './visibilityFilter'
import releases from './releases'

const reducerFilter = combineReducers({
  releases,
  visibilityFilter
})

export default reducerFilter
