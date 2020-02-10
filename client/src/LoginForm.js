import React from 'react';
import BaseForm from './BaseForm';
import { reduxForm } from 'redux-form';
import { login } from './actions/auth';

const LoginForm = (props) => {

	return (
		<BaseForm {...props} buttonText={'Login'} submitRequest={(values) => props.dispatch(login(values))}/>
	);
}

export default reduxForm({
	form: 'login'
})(LoginForm);