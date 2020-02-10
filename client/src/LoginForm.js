import React from 'react';
import { Field, reduxForm, focus, clearSubmitErrors } from 'redux-form'
import { login } from './actions/auth';
import { toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import './Form.css';

class LoginForm extends React.Component {
	handleLogin = (values) => {
		return this.props.dispatch(login(values))
			.catch(e => {
				console.log(e);
				toast.error(e.errors._error, {
					className: 'toast'
				});
			});
	};

	render = () => {
		const { handleSubmit } = this.props;
		const { pristine, submitting } = this.props;
		const disabled = pristine || submitting;

		return (
			<form onSubmit={handleSubmit(values => this.handleLogin(values))}>
				<div className='input'>
					<label htmlFor='username'>Username</label>
					<Field component='input' disabled={submitting} name='username' type='text'/>
				</div>
				<div className='input'>
					<label htmlFor='password'>Password</label>
					<Field component='input' disabled={submitting} name='password' type='text'/>
				</div>
				<button disabled={disabled} type="submit">Login</button>
			</form>
		);
	}
}

export default reduxForm({
	form: 'login',
	onSubmitFail: (errors, dispatch) => dispatch(focus('login', 'username')),
	onChange: (values, dispatch, props) => {
		if (props.error !== null) {
			dispatch(clearSubmitErrors('login'));
		}
	}
})(LoginForm);