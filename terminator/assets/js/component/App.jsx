import React from 'react'

import Footer from './Footer'

import ReleaseList from './ReleaseList'
import VisibleReleaseList from '../containers/VisibleRelease'
import FetchRelease from '../containers/FetchRelease'


const App = () => (
  <div>
    <VisibleReleaseList />
    <Footer />
  </div>
)

export default App
