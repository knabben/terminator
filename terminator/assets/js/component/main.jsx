import React from 'react'

import axios from 'axios'
import uuid from 'uuid'

import Release from './release'


class Main extends React.Component {

    constructor(props) {
        super();

        this.onDelete = this.onDelete.bind(this);

        // Fetch items list
        axios.get('/items/').then( (response) => {
            this.setState({releases: response.data.releases})
        })
    };

    componentWillMount() {
        this.setState({releases: []})
    }

    onDelete(name) {
        axios.delete('/items/').then( (response) => {
            console.log(response)
        })
        this.setState({
            releases: this.state.releases.filter(
                release => release.name !== name
            )
        })
    }

    render() {
        const releases = this.state.releases;
        return (
            <div className="row">
                <div className="col-12">
                    <table className="table table-stripped">
                        <thead className="thead-inverse">
                        <tr>
                            <td>Name</td>
                            <td>Version</td>
                            <td>Namespace</td>
                            <td>LastDeploy</td>
                            <td>Delete</td>
                        </tr>
                        </thead>
                        <tbody>
                        {releases.map(({name, namespace, version, last_deploy}) =>
                            <Release key={uuid.v4()}
                                name={name}
                                namespace={namespace}
                                last_deploy={last_deploy}
                                version={version}
                                onDelete={() => this.onDelete(name)}/>
                        )}
                        </tbody>
                    </table>
                </div>
            </div>
        )
    }

}

export default Main;
