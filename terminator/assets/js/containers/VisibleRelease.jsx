import { connect } from 'react-redux'
import ReleaseList from '../component/ReleaseList'
import { RECEIVE_RELEASE } from '../actions/'


const mapStateToProps = (state) => ({
  releases: state.releases
})
const VisibleReleaseList = connect(mapStateToProps)(ReleaseList)

export default VisibleReleaseList
