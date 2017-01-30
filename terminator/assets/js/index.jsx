import React from 'react'
import ReactDOM from 'react-dom'
import { Router, Route, hashHistory } from 'react-router'

import Main from './component/main'

ReactDOM.render(
    <Router history={hashHistory}>
        <Route path="/" component={Main}/>
    </Router>
    ,document.getElementById('react-app')
)
