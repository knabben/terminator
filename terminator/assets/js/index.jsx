import React from 'react'

import { render } from 'react-dom'

import { Provider } from 'react-redux'
import { createStore } from 'redux'

import reducerFilter from './reducers/index'

import App from './component/App'

let store = createStore(reducerFilter)

render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('app')
)
