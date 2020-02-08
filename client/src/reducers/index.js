import { combineReducers } from 'redux';
import albumReducer from './album';
import authReducer from './auth';
import { reducer as formReducer } from 'redux-form';

export default combineReducers({
    album: albumReducer,
    auth: authReducer,
    form: formReducer
});