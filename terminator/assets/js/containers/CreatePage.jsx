import React from 'react';
import { connect } from 'react-redux'
import { Link } from 'react-router'


class CreatePage extends React.Component {
  render() {
    console.log(this.props)
    return (
      <div>
        <br />
        <br />
        <h5>Creating charts here</h5>
        <Link to="/">Back</Link>
      </div>
    )
  }
}

export default connect()(CreatePage)