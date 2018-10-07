import React from 'react';
import styled from 'styled-components';

import messages from './messages';

import { connect } from 'react-redux';
import { compose } from 'redux';
import { createStructuredSelector } from 'reselect';

import injectReducer from 'utils/injectReducer';
import injectSaga from 'utils/injectSaga';

import reducer from './reducer';
import saga from './saga';

import {
  makeSelectVersion,
  makeSelectKind,
  makeSelectSpec,
  makeSelectStatus
} from './selectors';

import {
  sendDeleteCRD,
  sendCreateCRD
} from './actions';


const Header = styled.div`
  font-family: "term";
  text-align: center;
  color: #1d6686;
  font-size: 25px;
  padding: 1.2em 0 1.2em 0;
`;

const ContainerWrapper = styled.div`
  display: flex;
  height: 80px;
  padding-bottom: 1em;
  padding-top: 1em;
`;

const Label = styled.div`
  margin:auto;
  align-self: auto;
  flex: 1 100px;
`

const Pods = styled.div`
  align-self: auto;
  flex: 1 10px;
  margin:auto;
`;

const IconLay = styled.div`
  align-self: auto;
  flex: 1 60px;
  margin: auto;
`;

const Icon = styled.img`
  width: 50px;
  height: 50px;
`;

const Status = styled.div`
  align-self: auto;
  font-size: 10px;
  flex: 1 470px;
  margin: auto;
`;


export class Item extends React.PureComponent {
  render() {
    const { name, exists, status } = this.props;
    const iconPath = `icon-${name}.png`;

    return (
      <ContainerWrapper>
        <IconLay> <Icon src={iconPath} /> </IconLay>
        <Label> {name} </Label>
        <Status> {status} </Status>
        <Pods>
        {
          exists && <button className="button" onClick={() => this.props.onDelete(name)}>DELETE</button>
        }
        {
          !exists && <button onClick={() => this.props.onCreate(name)}>CREATE</button>
        }
        </Pods>
      </ContainerWrapper>
    );
  }
}

/* eslint-disable-line react/prefer-stateless-function */
export class HomePage extends React.PureComponent {
  render() {
    const { spec, status } = this.props;
    return (
      <div>
        <Header>TERMINATOR BACKING SERVICE</Header>
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
        <Item
          name="rabbitmq"
          exists={spec.rabbitmq}
          status={status.rabbitNode}
          onDelete={this.props.onDelete}
          onCreate={this.props.onCreate}
        />
        <Item
          name="elastic"
          exists={spec.elastic}
          status={status.elastic}
          onDelete={this.props.onDelete}
          onCreate={this.props.onCreate}
        />
      </div>
    );
  }
}

HomePage.defaultProps = {
  'spec': {
    'memcache': false
  },
  'status': {
    'memcacheNode': ''
  }
};

const mapStateToProps = createStructuredSelector({
  kind: makeSelectKind(),
  version: makeSelectVersion(),
  spec: makeSelectSpec(),
  status: makeSelectStatus()
});

export function mapDispatchToProps(dispatch) {
  return {
    onDelete: evt => dispatch(sendDeleteCRD(evt)),
    onCreate: evt => dispatch(sendCreateCRD(evt))
  };
}

const withConnect = connect(mapStateToProps, mapDispatchToProps);
const withReducer = injectReducer({ key: 'terminator', reducer });
const withSaga = injectSaga({ key: 'terminator', saga });

export default compose(
  withReducer,
  withSaga,
  withConnect,
)(HomePage);
