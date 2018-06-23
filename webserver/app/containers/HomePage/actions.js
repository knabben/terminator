import {
  TERMINATOR_PAYLOAD,
  DELETE_CRD
} from './constants';

export function sendDeleteCRD(name) {
  return {
    type: DELETE_CRD,
    name,
  };
}

export function sendTerminatorPayload(payload) {
  return {
    type: TERMINATOR_PAYLOAD,
    payload,
  };
}

