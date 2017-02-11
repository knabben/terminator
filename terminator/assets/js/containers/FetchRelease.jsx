import React from 'react'

import { connect } from 'react-redux'
import { addRelease, fetchReleases } from '../actions'


class FetchRelease extends React.Component {

  // Initialize async ITEMs request
  componentDidMount() {
    const { dispatch } = this.props
    dispatch(fetchReleases())
  }

  render() {
    return (
      <div></div>
    )
  }
}

FetchRelease = connect()(FetchRelease)


export default FetchRelease
