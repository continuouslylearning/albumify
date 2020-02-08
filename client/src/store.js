import { applyMiddleware, createStore } from 'redux';
import thunk from 'redux-thunk';
import combinedReducer from './reducers';
import { setAuthToken } from './actions/auth';
import { loadAuthToken } from './utils';

const store = createStore(combinedReducer, applyMiddleware(thunk));
const authToken = loadAuthToken();

if (authToken !== null) {
	const token = authToken;
	store.dispatch(setAuthToken(token));
}

export default store;