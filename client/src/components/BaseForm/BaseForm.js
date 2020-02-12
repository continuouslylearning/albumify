import React from 'react';
import { Field } from 'redux-form'
import { toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import './Form.css';

export default class extends React.Component {
	onSubmit = (values) => {
		return this.props.submitRequest(values)
			.catch(e => {
				toast.error(e.errors._error, {
					className: 'toast'
				});
			});
	};

	render = () => {
		const { buttonText } = this.props;
		const { handleSubmit, pristine, submitting } = this.props;
		const disabled = pristine || submitting;

		return (
			<form onSubmit={handleSubmit(values => this.onSubmit(values))}>
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
}