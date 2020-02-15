import React from 'react';
import { useDispatch } from 'react-redux';
import { clearAuth } from '../../actions/auth';
import logout from '../../images/logout.svg';
import './Header.css';

export default () => {
	const dispatch = useDispatch();

	return (
		<header>
			<div className="header-items">
				<h1>Albumify</h1>
				<img alt={logout} onClick={() => dispatch(clearAuth())} src={logout}/>
			</div>
		</header>
	);
};