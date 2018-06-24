/**
 * Gets the repositories of the user from Github
 */

import { take, all } from 'redux-saga/effects';
import { call, put, select, takeLatest } from 'redux-saga/effects';
import { DELETE_CRD } from './constants'

import { sendTerminatorPayload, repoLoadingError } from './actions';

import { eventChannel } from 'redux-saga'
import request from 'utils/request';

function* createSocketChannel(ws) {
  return eventChannel(emitter => {
    ws.onopen = () => {
      console.log('opening...')
    }

    ws.onerror = (error) => {
      console.log('WebSocket error ' + error)
      console.dir(error)
    }

    ws.onmessage = (e) => {
      console.log(e)
      try {
        const payload = JSON.parse(e.data).message
        emitter(JSON.parse(payload))
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

function* sendDeleteCRD(action, ws) {
  return eventChannel(emitter => {
    ws.onopen = () => {
      console.log("opening for send")
      emitter.send({'delete': action.name})
    }

    ws.onerror = (error) => {
      console.log('WebSocket error ' + error)
      console.dir(error)
    }

    const unsubscribe = () => {
      console.log('Socket off')
    }

    return unsubscribe
  })
}

export default function* wsSagas() {
  const wsUrl = "ws://localhost:8092/ws/events/"

  const ws = new WebSocket(wsUrl)
  const socket = yield call(createSocketChannel, ws)

  yield takeLatest(DELETE_CRD, sendDeleteCRD, ws)

  while (true) {
    const payload = yield take(socket)
    yield put(sendTerminatorPayload(payload))
  }
}
