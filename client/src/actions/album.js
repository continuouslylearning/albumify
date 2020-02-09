import axios from "axios";
import { SERVER_URI } from '../config';

export const GET_ALBUM_REQUEST = 'GET_ALBUM_REQUEST';
export const getAlbumRequest = () => {
	return {
		type: GET_ALBUM_REQUEST
	};
};

export const GET_ALBUM_SUCCESS = 'GET_ALBUM_SUCCESS';
export const getAlbumSuccess = (data) => {
	return {
		type: GET_ALBUM_SUCCESS,
		album: data
	};
};

export const GET_ALBUM_ERROR = 'GET_ALBUM_ERROR';
export const getAlbumError = (error) => {
	return {
		type: GET_ALBUM_ERROR,
		error
	};
};

export const getAlbum = () => async (dispatch, getState) => {
	dispatch(getAlbumRequest());
	try {
		const authToken = getState().auth.authToken;
		const config = {
			headers: {
				'Authorization': `Bearer ${authToken}`
			}
		};
		const { data } = await axios.get(SERVER_URI + '/album/', config);
		dispatch(getAlbumSuccess(data));
	} catch(e) {
		dispatch(getAlbumError(e));
	}
};

export const POST_IMAGES_REQUEST = 'POST_IMAGES_REQUEST';
export const postImagesRequest = () => {
	return {
		type: POST_IMAGES_REQUEST
	};
}

export const POST_IMAGES_SUCCESS = 'POST_IMAGES_SUCCESS';
export const postImagesSuccess = (images) => {
	return {
		images,
		type: POST_IMAGES_SUCCESS
	};
}

export const POST_IMAGES_ERROR = 'POST_IMAGES_ERROR';
export const postImagesError = (error) => {
	return {
		error,
		type: POST_IMAGES_ERROR
	};
}

export const postImages = (images) => async (dispatch, getState) => {
	dispatch(postImagesRequest());

	try {
		const formData = new FormData();
		const authToken = getState().auth.authToken;
		
		for (let i = 0; i < images.length; i++) {
			const image = images[i];
			formData.append(image.name, image);
		}
		
		const config = {
			headers: {
				'Authorization': `Bearer ${authToken}`,
				'Content-Type': 'multipart/form-data',
			}
		};
		const { data } = await axios.post(SERVER_URI + '/album/', formData, config);
		
		dispatch(postImagesSuccess(data));
	} catch(e) {
		dispatch(e);
	}
}
