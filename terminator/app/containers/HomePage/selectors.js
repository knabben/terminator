import { createSelector } from 'reselect';

const selectTerm = (state) => state.get('terminator');

const makeSelectVersion = () => createSelector(
  selectTerm,
  (termState) => termState.get('version')
);

const makeSelectKind = () => createSelector(
  selectTerm,
  (termState) => termState.get('kind')
);


const makeSelectStatus = () => createSelector(
  selectTerm,
  (termState) => termState.get('status')
);

const makeSelectSpec = () => createSelector(
  selectTerm,
  (termState) => termState.get('spec')
);

export {
  selectTerm,
  makeSelectVersion,
  makeSelectKind,
  makeSelectSpec,
  makeSelectStatus,
};
