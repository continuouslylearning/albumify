import React from 'react';
import { connect } from 'react-redux';
import LoginForm from './LoginForm';
import { ToastContainer } from 'react-toastify';
import './Landing.css';
import SignUpForm from './SignUpForm';

class Landing extends React.Component {
	constructor(props) {
		super(props);
	
		this.state = {
			displayed: 'login'
		};
	}

	toggleForm = (form) => {
		this.setState({
			displayed: form
		});
	}

	render = () => {
		const { displayed } = this.state;

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
									<span onClick={() => this.toggleForm('signUp')}>Sign up for an account</span>
								</>
							) : (
								<>
									<SignUpForm/>
									<span onClick={() => this.toggleForm('login')}>Already have an account?</span>
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
}

export default connect()(Landing);