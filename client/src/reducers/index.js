import { combineReducers } from 'redux';
import albumReducer from './album';
import { reducer as formReducer } from 'redux-form';

export default combineReducers({
    album: albumReducer,
    form: formReducer
});