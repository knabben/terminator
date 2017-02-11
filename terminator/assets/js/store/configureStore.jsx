import { createStore, applyMiddleware, compose } from 'redux'
import { persistState } from 'redux-devtools'

import DevTools from '../containers/DevTools'
import thunkMiddleware from 'redux-thunk'
import reducerFilter from '../reducers'
import createLogger from 'redux-logger'


const loggerMiddleware = createLogger()

const createStoreWithMiddleware = compose(
  applyMiddleware(
    thunkMiddleware
  ),
  DevTools.instrument(),
  persistState(getDebugSessionKey()),
)(createStore)

function getDebugSessionKey() {
  const matches = window.location.href.match(/[?&]debug_session=([^&]+)\b/);
  return (matches && matches.length > 0) ? matches[1] : null;
}

export default function configureStore(initialState) {
  const store = createStoreWithMiddleware(reducerFilter, initialState)
  if (module.hot) {
    module.hot.accept('../reducers', () => {
      const nextReducer = require('../reducers/index').default;
      store.replaceReducer(nextReducer);
    });
  }
  return store;
}
