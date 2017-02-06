import React from 'react'


// Presentional component, single item
class Release extends React.Component {
    render() {
      return (
        <tr>
          <td>{this.props.name}</td>
          <td>{this.props.version}</td>
          <td>{this.props.namespace}</td>
          <td>{this.props.last_deploy}</td>
        </tr>
      )
    }
}

Release.propTypes = {
  name: React.PropTypes.string,
  namespace: React.PropTypes.string,
  version: React.PropTypes.number,
}

export default Release
