import { take, all } from 'redux-saga/effects';
import { call, put, select, takeLatest } from 'redux-saga/effects';
import { DELETE_CRD } from './constants';

import { sendTerminatorPayload } from './actions';

import { eventChannel } from 'redux-saga';
import request from 'utils/request';

const wsUrl = "ws://localhost:8092/ws/events/";

function* createSocketChannel(ws) {
  return eventChannel(emitter => {
    ws.onopen = () => {
      console.log('Connected on websocket.');
    };

    ws.onerror = (error) => {
      console.log('websocket error ' + error);
      console.dir(error);
    };

    ws.onmessage = (e) => {
      try {
        const payload = JSON.parse(e.data).message;
        const data = JSON.parse(payload);
        console.log(data);
        emitter(data);
      } catch(err) {
        console.error(err);
      }
    };

    const unsubscribe = () => {
      console.log('socket off');
    };
    return unsubscribe;
  });
}

export default function* wsSagas() {
  const ws = new WebSocket(wsUrl);
  const socket = yield call(createSocketChannel, ws);
  while (true) {
    const payload = yield take(socket);
    yield put(sendTerminatorPayload(payload));
  }
}
