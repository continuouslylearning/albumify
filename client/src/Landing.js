import React from 'react';
import { connect } from 'react-redux';
import LoginForm from './LoginForm';
import { ToastContainer } from 'react-toastify';
import './Landing.css';

class Landing extends React.Component {
	render = () => {
		return (
			<div className='landing'>
				<div className='left'>
				</div>
				<div className='right'>
					<div className='header'>
						<h1>Albumify</h1>
					</div>
					<div className='form'>
						<LoginForm/>
					</div>
					<ToastContainer
						position="top-right"
						autoClose={5000}
						hideProgressBar={false}
						newestOnTop={false}
						closeOnClick
						rtl={false}
						pauseOnVisibilityChange
						draggable
						pauseOnHover
					/>
				</div>
			</div>
		);
	}
}

export default connect()(Landing);