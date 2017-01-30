import React from 'react'

class Release extends React.Component {
    render() {
        return (
            <div>
                {this.props.name} {this.props.namespace}
                <button onClick={this.props.onDelete}>x</button>
            </div>
        )
    }
}

Release.propTypes = {
    name: React.PropTypes.string,
    namespace: React.PropTypes.string
}

export default Release;
