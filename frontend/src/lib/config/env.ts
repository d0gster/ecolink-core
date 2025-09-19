interface AppConfig {
	apiUrl: string;
	googleClientId: string;
}

function validateConfig(): AppConfig {
	const apiUrl = import.meta.env.VITE_API_URL;
	const googleClientId = import.meta.env.VITE_GOOGLE_CLIENT_ID;

	if (!apiUrl) {
		throw new Error('VITE_API_URL environment variable is required');
	}

	if (!googleClientId) {
		throw new Error('VITE_GOOGLE_CLIENT_ID environment variable is required');
	}

	return {
		apiUrl,
		googleClientId
	};
}

export const config = validateConfig();