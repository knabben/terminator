import React from 'react'

import axios from 'axios'
import uuid from 'uuid'

import Release from './release'
import ListItem from './list'


class Main extends React.Component {

    constructor(props) {
        super();

        this.onDelete = this.onDelete.bind(this);

        // Fetch items list
        axios.get('/items/').then( (response) => {
            console.log(response.data.releases)
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
                <div className="col-12">{
                releases.map( ({name, namespace}) =>
                    <Release key={uuid.v4()}
                        name={name}
                        namespace={namespace}
                        onDelete={() => this.onDelete(name)}/>
                )}
                </div>
            </div>
        )
    }

}

export default Main;
