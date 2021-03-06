export const clearAuthToken = () => {
	try {
		localStorage.removeItem('authToken');
	} catch (e) { 
		return;
	}
};

export const loadAuthToken = () => {
	return localStorage.getItem('authToken');
};

export const saveAuthToken = (authToken) => {
	try {
		localStorage.setItem('authToken', authToken);
	} catch (e) { 
		return;
	}
};

export const getFileName = (URL) => {
	return URL.split('/').pop().split('?')[0];
};
