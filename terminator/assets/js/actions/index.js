import fetch from 'isomorphic-fetch'


export const REQUEST_RELEASE = 'REQUEST_RELEASE'
function requestRelease(release) {
  return {
    type: REQUEST_RELEASE,
    release
  }
}

export const RECEIVE_RELEASE = 'RECEIVE_RELEASE'
function receiveRelease(data) {
  return {
    type: RECEIVE_RELEASE,
    receiveAt: Date.now(),
    data
  }
}

export function fetchReleases() {
  // Fetch HELM releases async
  return dispatch => {
    return fetch("/items/").then(response => response.json())
      .then(data => dispatch(receiveRelease(data.releases)))
  }
}
