import React from 'react';
import { Field } from 'redux-form'
import { toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import './Form.css';

export default (props) => {
	const onSubmit = (values) => {
		return props.submitRequest(values)
			.catch(e => {
				toast.error(e.errors._error, {
					className: 'toast'
				});
			});
	};

	const { buttonText, handleSubmit, pristine, submitting } = props;
	const disabled = pristine || submitting;

	return (
		<form onSubmit={handleSubmit(values => onSubmit(values))}>
			<div className='input'>
				<label htmlFor='username'>Username</label>
				<Field component='input' disabled={submitting} name='username' type='text'/>
			</div>
			<div className='input'>
				<label htmlFor='password'>Password</label>
				<Field component='input' disabled={submitting} name='password' type='text'/>
			</div>
			<button disabled={disabled} type="submit">{buttonText}</button>
		</form>
	);
}