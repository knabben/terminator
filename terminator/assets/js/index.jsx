import React from 'react'

import { render } from 'react-dom'

import { Provider } from 'react-redux'
import DevTools from './containers/DevTools'
import configureStore from './store/configureStore'

import App from './component/App'
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
