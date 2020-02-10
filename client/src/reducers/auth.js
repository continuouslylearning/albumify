import { AUTH_ERROR, AUTH_REQUEST, AUTH_SUCCESS, CLEAR_AUTH, SET_AUTH_TOKEN } from '../actions/auth';
import { clearAuthToken } from '../utils';

const initialState = {
	authToken: null,
	error: null,
	loading: false,
	user: null
};

export default (state = initialState, action) => {
	if (action.type === AUTH_ERROR) {
		return {
			...state,
			error: action.error,
			loading: false
		};
	} else if (action.type === AUTH_REQUEST) {
		return {
			...state,
			error: null,
			loading: true
		};
	} else if (action.type === AUTH_SUCCESS) {
		return {
			...state,
			loading: false,
			user: action.user
		};
	} else if (action.type === CLEAR_AUTH) {
		clearAuthToken();
		
		return {
			...state,
			authToken: null,
			user: null
		};
	} else if (action.type ===  SET_AUTH_TOKEN) {
		return {
			...state,
			authToken: action.authToken
		};
	}
	return state;
};