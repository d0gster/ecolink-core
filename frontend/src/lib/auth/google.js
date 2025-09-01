// Google OAuth manual implementation
const GOOGLE_CLIENT_ID = import.meta.env.VITE_GOOGLE_CLIENT_ID;
const REDIRECT_URI = 'http://localhost:5173/auth/callback';

export function getGoogleAuthUrl() {
	const params = new URLSearchParams({
		client_id: GOOGLE_CLIENT_ID,
		redirect_uri: REDIRECT_URI,
		response_type: 'code',
		scope: 'openid email profile',
		access_type: 'offline',
		prompt: 'consent'
	});
	
	return `https://accounts.google.com/oauth/authorize?${params.toString()}`;
}

export async function exchangeCodeForTokens(code) {
	const response = await fetch('https://oauth2.googleapis.com/token', {
		method: 'POST',
		headers: {
			'Content-Type': 'application/x-www-form-urlencoded',
		},
		body: new URLSearchParams({
			client_id: GOOGLE_CLIENT_ID,
			client_secret: import.meta.env.VITE_GOOGLE_CLIENT_SECRET,
			code,
			grant_type: 'authorization_code',
			redirect_uri: REDIRECT_URI,
		}),
	});

	return response.json();
}

export async function getUserInfo(accessToken) {
	const response = await fetch('https://www.googleapis.com/oauth2/v2/userinfo', {
		headers: {
			Authorization: `Bearer ${accessToken}`,
		},
	});

	return response.json();
}