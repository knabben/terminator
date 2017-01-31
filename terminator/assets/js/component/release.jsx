import React from 'react'

class Release extends React.Component {
    render() {
        return (
            <tr>
                <td>{this.props.name}</td>
                <td>{this.props.version}</td>
                <td>{this.props.namespace}</td>
                <td>{this.props.last_deploy}</td>
                <td><button className="btn" onClick={this.props.onDelete}>x</button></td>
            </tr>
        )
    }
}

Release.propTypes = {
    name: React.PropTypes.string,
    namespace: React.PropTypes.string
}

export default Release;
