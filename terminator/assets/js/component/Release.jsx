import React from 'react'

// Presentional component, single item release
// rendering
class Release extends React.Component {
  render() {
    return (
      <div className="app-item">
        <div className="row">
          <div className="col-6 text-left">
            <span> {this.props.name} </span>
          </div>
          <div className="col-6">
            <div className="right-border text-right">
              version {this.props.version} | namespace {this.props.namespace}
            <button className="btn btn-danger btn-sm btn-del" onClick={this.props.onDelete}>x</button></div>
          </div>
        </div>
      </div>
    )
  }
}

Release.propTypes = {
  name: React.PropTypes.string,
  namespace: React.PropTypes.string,
  version: React.PropTypes.number,
}

export default Release
