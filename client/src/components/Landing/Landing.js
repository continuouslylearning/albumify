import React, { useState } from 'react';
import { ToastContainer } from 'react-toastify';
import LoginForm from '../LoginForm/LoginForm';
import SignUpForm from '../SignUpForm/SignUpForm';
import './Landing.css';

export default () => {
	const [ displayed, setDisplayed ] = useState('login');

	return (
		<div className='landing'>
			<div className='left'>
			</div>
			<div className='right'>
				<div className='header'>
					<h1>Albumify</h1>
				</div>
				<div className='form'>
					{
						displayed === 'login' ?
						(
							<>
								<LoginForm/>
								<span onClick={() => setDisplayed('signUp')}>Sign up for an account</span>
							</>
						) : (
							<>
								<SignUpForm/>
								<span onClick={() => setDisplayed('login')}>Already have an account?</span>
							</>
						)
					}
				</div>
				<ToastContainer
					autoClose={5000}
					closeOnClick
					draggable
					hideProgressBar={false}
					newestOnTop={false}
					pauseOnHover
					pauseOnVisibilityChange
					position="top-right"
					rtl={false}
				/>
			</div>
		</div>
	);
}
