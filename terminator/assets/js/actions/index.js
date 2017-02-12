import fetch from 'isomorphic-fetch';


export const REQUEST_RELEASE = 'REQUEST_RELEASE';
function requestRelease() {
  return {
      type: REQUEST_RELEASE,
      isFetching: true
  };
}

export const RECEIVE_RELEASE = 'RECEIVE_RELEASE';
function receiveRelease(data) {
  return {
    type: RECEIVE_RELEASE,
    receiveAt: Date.now(),
    data
  };
}

export const DELETE_RELEASE = 'DELETE_RELEASE';
function deleteRelease(releaseId) {
  return {
    type: DELETE_RELEASE,
    releaseId
  };
}


export const delRelease = releaseId => (dispatch, getState) => {
    dispatch(deleteReleaseId(releaseId));
};

// Fetch HELM releases async
export const fetchReleases = () => dispatch => {
    dispatch(requestRelease());
    return fetch("/items/")
        .then(response => response.json())
        .then(data => dispatch(receiveRelease(data.releases)));
};
