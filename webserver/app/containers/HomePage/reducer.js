/*
 * TerminatorReducer
 *
 */
import { fromJS } from 'immutable';

import {
  TERMINATOR_PAYLOAD,
} from './constants';

const initialState = fromJS({
  kind: undefined,
  version: undefined,
  spec: {
    memcache: false,
    redis: false,
  },
  status: {
    memcacheNode: [],
    redisNode: [],
  }
});

function terminatorReducer(state = initialState, action) {
  switch (action.type) {
  case TERMINATOR_PAYLOAD:
    return state
      .set('version', action.payload.apiVersion)
      .set('kind', action.payload.kind)
      .set('metadata', action.payload.metadata)
      .set('spec', action.payload.spec)
      .set('status', action.payload.status)
  default:
    return state;
  }
}

export default terminatorReducer;
