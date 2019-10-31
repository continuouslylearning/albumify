import { 
    GET_ALBUM_ERROR, 
    GET_ALBUM_REQUEST, 
    GET_ALBUM_SUCCESS
} from '../actions/album';

const initialState = {
    error: null,
    fetching: false,
    album: []
};

export default (state = initialState, action) => {
    if (action.type === GET_ALBUM_ERROR) {
        return {
            ...state,
            error: action.error,
            fetching: false,
            album: []
        };
    } else if (action.type === GET_ALBUM_REQUEST) {
        return {
            ...state,
            error: null,
            fetching: true
        };
    } else if (action.type === GET_ALBUM_SUCCESS) {
        return {
            ...state,
            fetching: false,
            album: action.album
        };
    }

    return state;
};