import { applyMiddleware, createStore } from 'redux';
import thunk from 'redux-thunk';
import combinedReducer from './reducers';

export default createStore(combinedReducer, applyMiddleware(thunk));