import React from 'react';
import BaseForm from '../BaseForm/BaseForm';
import { reduxForm } from 'redux-form';
import { register } from '../../actions/auth';

const SignUpForm = (props) => {
	return (
		<BaseForm {...props} buttonText={'Signup'} submitRequest={(values) => props.dispatch(register(values))}/>
	);
}

export default reduxForm({
	form: 'signUp'
})(SignUpForm);