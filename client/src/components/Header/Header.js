import React from 'react';
import { connect } from 'react-redux';
import { clearAuth } from '../../actions/auth';
import logout from '../../images/logout.svg';
import './Header.css';

const Header = (props) => {
	return (
		<header>
			<div className="header-items">
				<h1>Albumify</h1>
				<img alt={logout} onClick={() => props.logout()} src={logout}/>
			</div>
		</header>
	);
};

const mapDispatchToProps = (dispatch) => {
	return {
		logout: () => dispatch(clearAuth())
	};
};

export default connect(null, mapDispatchToProps)(Header);