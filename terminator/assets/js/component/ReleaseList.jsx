import React, { PropTypes } from 'react'
import uuid from 'uuid'

import Release from './Release'


class ReleaseList extends React.Component {
    render() {
        const { isFetching, data } = this.props
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
                        version={version} />
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
