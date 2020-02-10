
import jwtDecode from 'jwt-decode';
import axios from 'axios';
import { SubmissionError } from 'redux-form';
import { SERVER_URI } from '../config';
import { saveAuthToken } from '../utils';

export const AUTH_ERROR = 'AUTH_ERROR';
export const authError = error => ({
	type: AUTH_ERROR,
	error
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

export const CLEAR_AUTH = 'CLEAR_AUTH';
export const clearAuth = () => ({
	type: CLEAR_AUTH
});

export const SET_AUTH_TOKEN = 'SET_AUTH_TOKEN';
export const setAuthToken = authToken => ({
	type: SET_AUTH_TOKEN,
	authToken
});

const storeAuthInfo = (authToken, dispatch) => {
	const decodedToken = jwtDecode(authToken);
	dispatch(setAuthToken(authToken));
	dispatch(authSuccess(decodedToken.username));
	saveAuthToken(authToken);
};

export const login = ({ username, password }) => async (dispatch, getState) => {
	try {
		const res = await axios({
			method: 'POST',
			url: `${SERVER_URI}/login`,
			data: { username, password },
			headers: { 'Content-Type': 'application/json' }
		});
		
		storeAuthInfo(res.data.authToken, dispatch);
	} catch (e) {
		let message;
		if (!e.response || e.response.status === 500) {
			message = 'Unable to login. Please try again later.';
		} else {
			message = e.response.data.message;
		}

		throw new SubmissionError({
				_error: message
		});
	}
};

export const register = ({ username, password }) => async (dispatch, getState) => {
	try {
		const res = await axios({
			method: 'POST',
			url: `${SERVER_URI}/users`,
			data: { username, password },
			headers: { 'Content-Type': 'application/json' }
		});

		storeAuthInfo(res.data.authToken, dispatch);
	} catch(e) {
		let message;
		if (!e.response || e.response.status === 500) {
			message = 'Unable to login. Please try again later.';
		} else {
			message = e.response.data.message;
		}

		throw new SubmissionError({
				_error: message
		});
	}
};