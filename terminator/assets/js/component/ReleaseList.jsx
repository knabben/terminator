import React from 'react'
import uuid from 'uuid'

import Release from './Release'


const ReleaseList = ({ releases, onAddClick }) => (
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

// class ReleaseList extends React.Component {

//   componentWillMount() {
//     // Set initial state with fake data (getting from REST api after)
//     this.setState(
//       {"releases": [
//         //         {
//           "id": 2,
//           "version": 1,
//           "name": "dapper-hare",
//           "last_deploy": "2016-12-05T14:46:16",
//           "first_deploy": "2016-12-05T14:46:16",
//           "namespace": "axado"
//         },
//         {
//           "id": 3,
//           "version": 1,
//           "name": "sartorial-wasp",
//           "last_deploy": "2017-01-14T08:58:54",
//           "first_deploy": "2017-01-14T08:58:54",
//           "namespace": "monitoring"
//         }
//       ]
//     });
//   }

//   render() {
//     const releases = this.state.releases;
//     return (
//       <div className="row">
//         <div className="col-12">
//           <table className="table table-stripped">
//             <thead className="thead-inverse">
//               <tr>
//                 <td>Name</td>
//                 <td>Version</td>
//                 <td>Namespace</td>
//                 <td>LastDeploy</td>
//               </tr>
//               </thead>
//               <tbody>
//                 {releases.map(({name, namespace, version, last_deploy}) =>
//                   <Release key={uuid.v4()} name={name} namespace={namespace}
//                            last_deploy={last_deploy} version={version} />
//                 )}
//               </tbody>
//           </table>
//         </div>
//       </div>
//     )
//   }
// }

export default ReleaseList;
