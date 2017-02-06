const release = (state, action) => {
  switch(action.type) {
    case 'ADD_RELEASE':
      return action.release;
    default:
      return state
  }
}

const releases = (state = [], action) => {
  switch(action.type) {
    case 'ADD_RELEASE':
      return [
        ...state,
        release(undefined, action)
      ]
    default:
      return state
  }
}

export default releases
