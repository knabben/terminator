import { connect } from 'react-redux'
import ReleaseList from '../component/ReleaseList'

const getVisibleRelease = (releases, filter) => {
  switch (filter) {
    case 'SHOW_ALL':
      return releases
    case 'SHOW_NAMESPACE':
      return (
        releases.filter(r => r.namespace == 'monitoring')
      )
    default:
      throw new Error("Unknown filter:" + filter)
  }
}

const mapStateToProps = (state) => ({
  releases: getVisibleRelease(state.releases, state.visibilityFilter)
})

const VisibleReleaseList = connect(
  mapStateToProps
)(ReleaseList)

export default VisibleReleaseList
