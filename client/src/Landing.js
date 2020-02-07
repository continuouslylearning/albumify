import React from 'react';
import { connect } from 'react-redux';
import './Landing.css';

export default class Landing extends React.Component {
    render = () => {
        return (
            <div className='landing'>
                <div className='left'>
                </div>
                <div className='right'>
                    <h1>Albumify</h1>
                </div>
            </div>
        );
    }
}