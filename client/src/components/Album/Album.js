import React, { useEffect, useState} from 'react';
import Modal from 'react-modal';
import { CircleLoader as Loader } from "react-spinners";
import { useSelector, useDispatch } from 'react-redux';
import { deleteImage, getAlbum } from '../../actions/album';
import { getFileName } from '../../utils';
import './Album.css';
import dragAndDrop from '../../images/draganddrop.png';

Modal.setAppElement('body');

export default () => {
	const [ modalIsOpen, setModalIsOpen ] = useState(false);
	const [ openedImage, setOpenedImage ] = useState(null);
	const dispatch = useDispatch();

	const album = useSelector(state => state.album.album);
	const error = useSelector(state => state.album.error);
	const fetching = useSelector(state => state.album.fetching);

	useEffect(() => {
		dispatch(getAlbum())
	}, [dispatch]);

	const onClick = (e) => {
		setModalIsOpen(true);
		setOpenedImage(e.target.src);
	}

	const onKeyPress = (e) => {
		const imageSRC = e.target.querySelector('.content').src;
		const imageKey = getFileName(imageSRC);

		if (e.key === 'Delete') {
			return dispatch(deleteImage(imageKey))
				.then(() => dispatch(getAlbum()));
		}
	}

	if (error !== null) {
		return <div className="error">COULD NOT LOAD</div>;
	}

	if (fetching) {
		return (
			<div className='spinner'>
				<Loader loading={fetching}/>
			</div>
		);
	}

	if (album.length === 0) {
		return (
			<div className='drag-and-drop'>
				<img alt={dragAndDrop} src={dragAndDrop}/>
				<span>Drag your image files into the dashboard</span>
			</div>
		)
	}

	return (
		<div className="album-container">
			<div className="album">
				{album.map((image, index) => {
					return (
						<div className='thumbnail' key={image} onKeyDown={onKeyPress} tabIndex={index}> 
							<img alt={image} className='content' key={index} onClick={onClick} src={image}/>
						</div>
					);
				})}
			</div>
			<Modal 
				className="image-modal"
				isOpen={modalIsOpen} 
				onRequestClose={() => setModalIsOpen(false)}
				overlayClassName="modal-overlay"
			>
				<img src={openedImage} alt={openedImage}/>
			</Modal>
		</div>
	);
}