import { connect } from 'react-redux'
import ReleaseList from '../component/ReleaseList'
import { RECEIVE_RELEASE } from '../actions/'

function mapStateToProps(state) {
    const { releases } = state
    return {
        data: releases.data || [],
        isFetching: releases.isFetching
    }
}

const VisibleReleaseList = connect(mapStateToProps)(ReleaseList)

export default VisibleReleaseList
