import { RECEIVE_RELEASE, REQUEST_RELEASE } from '../actions/'


const releases = (state = [], action) => {
  switch(action.type) {

    case REQUEST_RELEASE:
      return Object.assign({}, state, {
          isFetching: true,
          data: []
      })

    case RECEIVE_RELEASE:
      return Object.assign({}, state, {
          data: action.data,
          isFetching: false
      })

    default:
      return state
  }
}

export default releases
