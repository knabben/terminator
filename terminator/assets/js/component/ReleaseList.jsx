import React from 'react'
import uuid from 'uuid'

import Release from './Release'


const ReleaseList = ({ releases }) => (
  <div className="row">
    <div className="col-12">
      <table className="table table-stripped">
        <thead className="thead-inverse">
          <tr>
            <td>Name</td>
            <td>Version</td>
            <td>Namespace</td>
            <td>LastDeploy</td>
          </tr>
          </thead>
          <tbody>
            {releases.map(({name, namespace, version, last_deploy}) =>
              <Release key={uuid.v4()} name={name} namespace={namespace}
                       last_deploy={last_deploy} version={version} />
            )}
          </tbody>

      </table>
    </div>
  </div>
)

export default ReleaseList;
