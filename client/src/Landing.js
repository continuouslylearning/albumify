import React from 'react';
import { connect } from 'react-redux';
import LoginForm from './LoginForm';
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
				</div>
			</div>
		);
	}
}

export default connect()(Landing);