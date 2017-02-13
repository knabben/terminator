import React from 'react'
import VisibleReleaseList from '../containers/VisibleRelease'

import { Link } from 'react-router'

const App = () => (
  <div>
    <div className="create-label">
      <Link to="/create">Create</Link>
    </div>
    <VisibleReleaseList />
  </div>
)

export default App
