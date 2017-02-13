import React , { PropTypes } from 'react';

import { render } from 'react-dom'
import { Router, Route, browserHistory } from 'react-router'
import { Provider } from 'react-redux';

import App from '../component/App'
import { fetchReleases } from '../actions/'

import DevTools from '../containers/DevTools'
import CreatePage from '../containers/CreatePage'


const Root = ({ store, history }) => (
  <Provider store={store}>
    <div>
      <Router history={history}>
        <Route path="/" component={App} />
        <Route path="/create" component={CreatePage} />
      </Router>
      <DevTools />
    </div>
  </Provider>
)

export default Root