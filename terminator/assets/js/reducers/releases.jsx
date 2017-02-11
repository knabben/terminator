import { RECEIVE_RELEASE } from '../actions/'

// Reducer for RELEASES
const releases = (state = [], action) => {
  switch(action.type) {
    case RECEIVE_RELEASE:
      return action.data
    default:
      return state
  }
}

export default releases
