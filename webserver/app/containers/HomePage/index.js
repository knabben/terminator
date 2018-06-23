import React from 'react';
import PropTypes from 'prop-types';
import { Helmet } from 'react-helmet';
import { FormattedMessage } from 'react-intl';
import { connect } from 'react-redux';
import { compose } from 'redux';
import { createStructuredSelector } from 'reselect';

import styled from 'styled-components';

import injectReducer from 'utils/injectReducer';
import injectSaga from 'utils/injectSaga';
import H2 from 'components/H2';
import A from 'components/A';
import ReposList from 'components/ReposList';
import AtPrefix from './AtPrefix';
import CenteredSection from './CenteredSection';
import Form from './Form';
import Input from './Input';
import Section from './Section';
import messages from './messages';
import { loadRepos } from '../App/actions';
import { makeSelectVersion, makeSelectKind, makeSelectSpec, makeSelectStatus } from './selectors';
import reducer from './reducer';
import saga from './saga';

import { sendDeleteCRD } from './actions'

const Header = styled.div`
text-align: center;
font-size: 12px;
`

export class Item extends React.PureComponent {
  render() {
    const { name, exists, status } = this.props

    return (
      <div>
        {name} - {exists && <button onClick={() => this.props.onDelete(name)}>delete</button>} - {status}
      </div>
    )
  }
}

export class HomePage extends React.PureComponent { // eslint-disable-line react/prefer-stateless-function
  render() {
    const { spec, status } = this.props;

    return (
      <div>
        <Item
            name="memcache"
            exists={spec.memcache}
            status={status.memcacheNode}
            onDelete={this.props.onDelete}
        />
        <Item
            name="redis"
            exists={spec.redis}
            status={status.redisNode}
            onDelete={this.props.onDelete}
        />
      </div>
    )
  }
}


HomePage.propTypes = {
  loading: PropTypes.bool,
  error: PropTypes.oneOfType([
    PropTypes.object,
    PropTypes.bool,
  ]),
  repos: PropTypes.oneOfType([
    PropTypes.array,
    PropTypes.bool,
  ]),
};

const mapStateToProps = createStructuredSelector({
  kind: makeSelectKind(),
  version: makeSelectVersion(),
  spec: makeSelectSpec(),
  status: makeSelectStatus(),
});

export function mapDispatchToProps(dispatch) {
  return {
    onDelete: evt => dispatch(sendDeleteCRD(evt))
  }
}

const withConnect = connect(mapStateToProps, mapDispatchToProps);
const withReducer = injectReducer({ key: 'terminator', reducer });
const withSaga = injectSaga({ key: 'terminator', saga });

export default compose(
  withReducer,
  withSaga,
  withConnect,
)(HomePage);
