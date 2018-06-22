/**
 * Gets the repositories of the user from Github
 */

import { take, all } from 'redux-saga/effects';
import { call, put, select, takeLatest } from 'redux-saga/effects';
import { LOAD_REPOS } from 'containers/App/constants';
import { reposLoaded, repoLoadingError } from 'containers/App/actions';

import { eventChannel } from 'redux-saga'
import request from 'utils/request';
import { makeSelectUsername } from 'containers/HomePage/selectors';

/**
 * Github repos request/response handler
 */
export function* getRepos() {
  // Select username from store
  const username = yield select(makeSelectUsername());
  const requestURL = `https://api.github.com/users/${username}/repos?type=all&sort=updated`;

  try {
    // Call our request helper (see 'utils/request')
    const repos = yield call(request, requestURL);
    yield put(reposLoaded(repos, username));
  } catch (err) {
    yield put(repoLoadingError(err));
  }
}

function createSocketChannel() {
  const wsUrl = "ws://localhost:8000/ws/events/"

  return eventChannel(emitter => {
    const ws = new WebSocket(wsUrl)

    ws.onopen = () => {
      console.log('opening...')
    }

    ws.onerror = (error) => {
      console.log('WebSocket error ' + error)
      console.dir(error)
    }


    ws.onmessage = (e) => {
      try {
        const payload = JSON.parse(e.data).message
        console.log(JSON.parse(payload))
      } catch(err) {
        console.error(err)
      }
    }

    // unsubscribe function
    const unsubscribe = () => {
      console.log('Socket off')
    }

    return unsubscribe
  })
}

export default function* wsSagas() {
  const socket = yield call(createSocketChannel)

  while (true) {
    const payload = yield take(socket)
  }
}
