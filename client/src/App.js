import React from 'react';
import { connect } from 'react-redux';
import AlbumContainer from './AlbumContainer';
import Header from './Header';
import Landing from './Landing';
import './App.css';

class App extends React.Component {
    render = () => {
        const { loggedIn } = this.props;

        if (loggedIn) {
            return (
                <>  
                    <Header/>
                    <AlbumContainer/>
                </>
            );
        }

        return (
            <>
                <Landing/>
            </>
        );
    };
}

const mapStateToProps = (state) => {
    const { authToken } = state.auth;

    return {
        loggedIn: authToken != null
    };
};

export default connect(mapStateToProps)(App);

