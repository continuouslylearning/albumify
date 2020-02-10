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
	let fileName;
	fileName = URL.split('/').pop();
	fileName = fileName.split('?')[0];

	return fileName;
};
