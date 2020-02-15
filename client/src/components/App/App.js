import React from 'react';
import { useSelector } from 'react-redux';
import AlbumContainer from '../AlbumContainer/AlbumContainer';
import Header from '../Header/Header';
import Landing from '../Landing/Landing';
import './App.css';

export default () => {
	const loggedIn = useSelector(state => state.auth.authToken != null);

	if (loggedIn) {
		return (
			<>  
				<Header/>
				<AlbumContainer/>
			</>
		);
	}

	return (
		<Landing/>
	);
}
