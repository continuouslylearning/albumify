import React from 'react';
import Modal from 'react-modal';
import { CircleLoader as Loader } from "react-spinners";
import { connect } from 'react-redux';
import { deleteImage, getAlbum } from '../../actions/album';
import { getFileName } from '../../utils';
import './Album.css';
import dragAndDrop from '../../images/draganddrop.png';

Modal.setAppElement('body');

class Album extends React.Component {
	constructor(props) {
		super(props);

		this.state = {
			modalIsOpen: false,
			openedImage: null
		};
	}

	componentDidMount = () => {
		return this.props.getAlbum();
	}

	closeModal = () => {
		this.setState({ 
			modalIsOpen: false
		});
	}

	onClick = (e) => {
		this.setState({
			modalIsOpen: true,
			openedImage: e.target.src
		});
	}

	onKeyPress = (e) => {
		const imageSRC = e.target.querySelector('.content').src;
		const imageKey = getFileName(imageSRC);

		if (e.key === 'Delete') {
			return this.props.deleteImage(imageKey)
				.then(() => this.props.getAlbum());
		}
	}

	render = () => {
		const { album, error, fetching } = this.props;
		const { modalIsOpen, openedImage } = this.state;

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
							<div className='thumbnail' key={image} onKeyDown={this.onKeyPress} tabIndex={index}> 
								<img alt={image} className='content' key={index} onClick={this.onClick} src={image}/>
							</div>
						);
					})}
				</div>
				<Modal 
					className="image-modal"
					isOpen={modalIsOpen} 
					onRequestClose={this.closeModal}
					overlayClassName="modal-overlay"
				>
					<img src={openedImage} alt={openedImage}/>
				</Modal>
			</div>
		);
	}
}

const mapStateToProps = (state) => {
	const { error, fetching, album } = state.album;

	return {
		album,
		error,
		fetching
	};
};

const mapDispatchToProps = (dispatch) => {
	return {
		deleteImage: (key) => dispatch(deleteImage(key)),
		getAlbum: () => dispatch(getAlbum())
	};
};

export default connect(mapStateToProps, mapDispatchToProps)(Album);