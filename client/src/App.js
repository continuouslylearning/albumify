import React from 'react';
import { connect } from 'react-redux';
import Album from './Album';
import Landing from './Landing';

class App extends React.Component {
    render = () => {
        const { loggedIn } = this.props;

        if (loggedIn) {
            return <Album/>;
        }

        return <Landing/>;
    };
}

const mapStateToProps = (state) => {
    const { user } = state.auth;

    return {
        loggedIn: user !== null
    };
};

export default connect(mapStateToProps)(App);

