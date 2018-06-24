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
import Button from 'components/Button';
import ReposList from 'components/ReposList';
import messages from './messages';

import reducer from './reducer';
import saga from './saga';
import { makeSelectVersion, makeSelectKind, makeSelectSpec, makeSelectStatus } from './selectors';
import { sendDeleteCRD, sendCreateCRD } from './actions'

const Header = styled.div`
  text-align: center;
  color: #41addd;
  font-size: 20px;
  padding: 1.2em 0 1.2em 0;
`

const ContainerWrapper = styled.div`
  display: flex;
  height: 80px;
  padding-bottom: 1em;
  padding-top: 1em;
`

const Label = styled.div`
  margin:auto;
  align-self: auto;
  flex: 1 100px;
`

const Pods = styled.div`
  align-self: auto;
  flex: 1 10px;
  margin:auto;
`

const IconLay = styled.div`
  align-self: auto;
  flex: 1 60px;
  margin: auto;
`

const Icon = styled.img`
  width: 50px;
  height: 50px;
`
const Status = styled.div`
  align-self: auto;
  font-size: 10px;
  flex: 1 470px;
  margin: auto;
`

export class Item extends React.PureComponent {
  render() {
    const { name, exists, status } = this.props
    const iconPath = `icon-${name}.png`

    return (
      <ContainerWrapper>
        <IconLay> <Icon src={iconPath} /> </IconLay>
        <Label> {name} </Label>
        <Status> {status} </Status>
        <Pods>
          {exists &&
           <Button onClick={() => this.props.onDelete(name)}>DELETE</Button>
          }

          {!exists &&
           <Button onClick={() => this.props.onCreate(name)}>CREATE</Button>
          }
        </Pods>
      </ContainerWrapper>
    )
  }
}


export class HomePage extends React.PureComponent { // eslint-disable-line react/prefer-stateless-function
  render() {
    const { spec, status } = this.props;

    return (
      <div>
        <Header>TBS - Terminator Backing Services - CONSOLE</Header>
        <Item
            name="memcache"
            exists={spec.memcache}
            status={status.memcacheNode}
            onDelete={this.props.onDelete}
            onCreate={this.props.onCreate}
        />
        <Item
            name="redis"
            exists={spec.redis}
            status={status.redisNode}
            onDelete={this.props.onDelete}
            onCreate={this.props.onCreate}
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
    onDelete: evt => dispatch(sendDeleteCRD(evt)),
    onCreate: evt => dispatch(sendCreateCRD(evt)),
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
