import React from 'react';
import { connect } from 'react-redux';
import { clearAuth } from './actions/auth';
import './Header.css';
import logout from './images/logout.svg';

const Header = (props) => {
    return (
        <header>
            <div className="header-items">
                <h1>Albumify</h1>
                <img alt={logout} onClick={() => props.dispatch(clearAuth())} src={logout}/>
            </div>
        </header>
    );
};

export default connect()(Header);