import {
  TERMINATOR_PAYLOAD,
  DELETE_CRD
} from './constants';

import request from 'utils/request';


export function sendDeleteCRD(name) {
  const requestURL = `http://localhost:8092/crd`;
  try {
    const options = {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({name: name})
    };
    const response = request(requestURL, options);
  } catch (err) {
    console.error(err);
  }
}

export function sendCreateCRD(name) {
  const requestURL = `http://localhost:8092/crd`;
  try {
    const options = {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({name: name})
    };
    const response = request(requestURL, options);
  } catch (err) {
    console.error(err);
  }
}

export function sendTerminatorPayload(payload) {
  return {
    type: TERMINATOR_PAYLOAD, payload
  };
}

