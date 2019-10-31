import { combineReducers } from 'redux';
import albumReducer from './album';

export default combineReducers({
    album: albumReducer
});