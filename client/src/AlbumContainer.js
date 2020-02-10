import DropZone from 'react-dropzone';
import React from 'react'
import { connect } from 'react-redux';
import { ToastContainer, toast } from 'react-toastify';

import Album from './Album';
import { getAlbum, postImages} from './actions/album';

const style = {
	height: 'auto',
	minHeight: '100%'
};

class AlbumContainer extends React.Component {
	onDrop = (files) => {
		const images = files.filter(file => {
			return /image.*/.test(file.type) && file.size < 2097152;
		});

		if (files.length > images.length) {
			toast.error('Only images under 2MB can be uploaded', {
				className: 'toast',
				position: toast.POSITION.BOTTOM_CENTER
			});
		}

		if (images.length === 0) {
			return;
		}

		return this.props.postImages(images)
			.then(() => this.props.getAlbum());
	}

	render = () => {
		return (
			<DropZone
				onDrop={this.onDrop}
				style={style}
			>   
				{ ({getRootProps, getInputProps}) => 
					(   
						<div {...getInputProps({ style })} {...getRootProps()}>
							<ToastContainer
								autoClose={5000}
								closeOnClick
								draggable
								hideProgressBar={false}
								newestOnTop={false}
								pauseOnHover
								pauseOnVisibilityChange
								position="top-right"
								rtl={false}
							/>
							<Album/>
						</div>
					)
				}   
			</DropZone>
		);
	}
}

const mapDispatchToProps = (dispatch) => {
	return {
		getAlbum: () => dispatch(getAlbum()),
		postImages: (images) => dispatch(postImages(images))
	};
}

export default connect(null, mapDispatchToProps)(AlbumContainer);