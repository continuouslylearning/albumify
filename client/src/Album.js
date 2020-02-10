import React from 'react';
import Modal from 'react-modal';
import { CircleLoader as Loader } from "react-spinners";
import { connect } from 'react-redux';
import { getAlbum } from './actions/album';
import './Album.css';

Modal.setAppElement('body');

class Album extends React.Component {
	constructor(props) {
		super(props);

		this.state = {
			modalisOpen: false,
			openedImage: null,
			isJpg: null
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
		const i = e.target.src.lastIndexOf(".");
		const fileExtension = e.target.src.slice(i+1, i+4);

		this.setState({
			modalIsOpen: true,
			openedImage: e.target.src,
			isJpg: fileExtension === 'jpg'
		});
	}

	render = () => {
		const { album, error, fetching } = this.props;
		const { modalIsOpen, openedImage, isJpg } = this.state;

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

		return (
			<div className="album-container">
				<div className="album">
					{album
						.map((image, index) => {
							const i = image.lastIndexOf(".");
							const fileExtension = image.slice(i+1, i+4);

							if (fileExtension === 'jpg') {
								return (
									<img alt={image} className='thumbnail' key={index} onClick={this.onClick} src={image}/>
								);
							}
							return (
								<video className='thumbnail' key={index} onClick={this.onClick} src={image}/>
							);
					})}
				</div>
				<Modal 
					className="image-modal"
					isOpen={modalIsOpen} 
					onRequestClose={this.closeModal}
					overlayClassName="modal-overlay"
				>
					{ isJpg 
						? <img src={openedImage} alt={openedImage}/>
						: (
							<video autoPlay controls loop src={openedImage} width='auto'>
							</video>
						)
					}
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
		getAlbum: () => dispatch(getAlbum())
	};
};

export default connect(mapStateToProps, mapDispatchToProps)(Album);