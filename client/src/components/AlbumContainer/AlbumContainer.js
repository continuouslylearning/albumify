import DropZone from 'react-dropzone';
import React from 'react'
import { useDispatch } from 'react-redux';
import { ToastContainer, toast } from 'react-toastify';
import Album from '../Album/Album';
import { getAlbum, postImages } from '../../actions/album';

const style = {
	height: 'auto',
	minHeight: '100%'
};

export default () => {
	const dispatch = useDispatch();
	const onDrop = (files) => {
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

		return dispatch(postImages(images))
			.then(() => dispatch(getAlbum()));
	}

	return (
		<DropZone
			onDrop={onDrop}
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

