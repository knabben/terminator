import React, { PropTypes } from 'react'
import uuid from 'uuid'

import { delRelease, fetchReleases } from '../actions/'
import Release from './Release'

class ReleaseList extends React.Component {
    constructor() {
      super()
      this.deleteRelease = this.deleteRelease.bind(this);
    }

    componentWillMount() {
        const { dispatch } = this.props
        dispatch(fetchReleases())
    } 

    deleteRelease(releaseName) {
      this.props.dispatch(delRelease(releaseName));
    }

    render() {
        const { isFetching, data, dispatch } = this.props
        return (
            <div>
              {isFetching &&
                <div className="loading text-center">
                  Loading...
                </div>
              }

              {!isFetching &&
                <div className="app-list">
                    {data.map(({name, namespace, version, last_deploy}) =>
                      <Release key={uuid.v4()}
                        name={name}
                        namespace={namespace}
                        last_deploy={last_deploy}
                        version={version}
                        onDelete={ () => this.deleteRelease(name) }
                        />
                    )}
                </div>
              }
            </div>
        )
    }
}

ReleaseList.propTypes = {
  isFetching: PropTypes.bool,
  data: PropTypes.arrayOf(PropTypes.shape({
    name: PropTypes.string.idRequired,
    namespace: PropTypes.string.idRequired,
    version: PropTypes.string.idRequired,
  })).isRequired
}

export default ReleaseList;
