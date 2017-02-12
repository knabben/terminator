require('./main.css');

import React from 'react'

import App from './component/App'
import DevTools from './containers/DevTools'
import configureStore from './store/configureStore'

import { render } from 'react-dom'
import { Provider } from 'react-redux'
import { fetchReleases } from './actions/'


let store = configureStore()

render(
  <Provider store={store}>
    <div>
      <App />
      <DevTools />
    </div>
  </Provider>,
  document.getElementById('app')
)

store.dispatch(fetchReleases())
