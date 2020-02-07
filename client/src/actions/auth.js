
import jwtDecode from 'jwt-decode';
import axios from 'axios';
import { SubmissionError } from 'redux-form';
import { SERVER_URI } from '../config';

export const SET_AUTH_TOKEN = 'SET_AUTH_TOKEN';
export const setAuthToken = authToken => ({
	type: SET_AUTH_TOKEN,
	authToken
});

export const CLEAR_AUTH = 'CLEAR_AUTH';
export const clearAuth = () => ({
	type: CLEAR_AUTH
});

export const AUTH_REQUEST = 'AUTH_REQUEST';
export const authRequest = () => ({
	type: AUTH_REQUEST
});

export const AUTH_SUCCESS = 'AUTH_SUCCESS';
export const authSuccess = user => ({
	type: AUTH_SUCCESS,
	user
});

export const AUTH_ERROR = 'AUTH_ERROR';
export const authError = error => ({
	type: AUTH_ERROR,
	error
});

const storeAuthInfo = (authToken, dispatch) => {
	const decodedToken = jwtDecode(authToken);
	// dispatch(setAuthToken(authToken));
	// dispatch(authSuccess(decodedToken.user));
	// saveAuthToken(authToken);
};

export const login = ({ username, password }) => (dispatch, getState) => {
	dispatch(authRequest());
    
    return axios({
		method: 'post',
		url: `${SERVER_URI}/auth/login`,
		headers: { 'Content-Type': 'application/json' },
		data: { username, password }
	})	
		.then(res => storeAuthInfo(res.data.authToken, dispatch))
		.catch(err => {
			// if server does not respond to the request, err object does not have a response property
			let message;
			if (!err.response || err.response.status === 500) {
				message = 'Unable to login. Please try again later.';
			} else {
				message = err.response.data.message;
			}

            // return Promise.reject(
            //     new SubmissionError({
            //         _error: message
            //     })
            // );

			dispatch(authError({ message }));
		});
};