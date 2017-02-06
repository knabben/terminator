import React from 'react'

import { connect } from 'react-redux'
import { addRelease } from '../actions'

let FetchRelease = ({ dispatch }) => {
  let release = {
    "id": 1,
    "version": 1,
    "name": "boiling-shrimp",
    "last_deploy": "2017-01-20T17:40:07",
    "first_deploy": "2017-01-20T17:40:07",
    "namespace": "monitoring"
  }

  console.log(release)
  return (
    <div className="col-12">
      <span onClick={() => dispatch(addRelease(release))}>add</span>
    </div>
  )
}

FetchRelease = connect()(FetchRelease)

export default FetchRelease
