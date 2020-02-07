import React from 'react';
import { Field, reduxForm } from 'redux-form'
import { login } from './actions/auth';
import './form.css';

class LoginForm extends React.Component {
	handleLogin = (values) => {
		this.props.dispatch(login(values));
	};

	render = () => {
		const { handleSubmit } = this.props;

		return (
			<form onSubmit={handleSubmit(values => this.handleLogin(values))}>
				<div className='input'>
					<label htmlFor='username'>Username</label>
					<Field name='username' component='input' type='text'/>
				</div>
				<div className='input'>
					<label htmlFor='password'>Password</label>
					<Field name='password' component='input' type='text'/>
				</div>
				<button type="submit">Login</button>
			</form>
		);
	}
}

export default reduxForm({
	form: 'login'
})(LoginForm);