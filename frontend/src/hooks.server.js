// Removendo Auth.js - implementaremos OAuth manual
export const handle = async ({ event, resolve }) => {
	return resolve(event);
};