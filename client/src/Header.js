import React from 'react';
import { connect } from 'react-redux';
import { clearAuth } from './actions/auth';
import './Header.css';

const Header = (props) => {
    return (
        <header>
            <h1>Albumify</h1>
            <button onClick={() => props.dispatch(clearAuth())}>Logout</button>
        </header>
    );
};

export default connect()(Header);