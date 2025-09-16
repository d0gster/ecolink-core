import { writable } from 'svelte/store';
import { browser } from '$app/environment';
import type { User } from '$lib/stores/auth.ts';

export const isLoading = writable(true);
export const isAuthenticated = writable(false);
export const user = writable<User | null>(null);

const GOOGLE_CLIENT_ID = import.meta.env.VITE_GOOGLE_CLIENT_ID;
const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080';
let REDIRECT_URI = 'http://localhost:5173/auth/callback/google';

if (browser && window) {
	REDIRECT_URI = `${window.location.origin}/auth/callback/google`;
}

export function initGoogleAuth() {
	if (!browser) return;
	
	// Check if saved user exists
	const savedUser = localStorage.getItem('ecolink_user');
	const sessionToken = localStorage.getItem('ecolink_session_token');
	
	if (savedUser && sessionToken) {
		try {
			const userData = JSON.parse(savedUser);
			user.set(userData);
			isAuthenticated.set(true);
		} catch (error) {
			console.error('Erro ao carregar usuário salvo:', error);
			localStorage.removeItem('ecolink_user');
			localStorage.removeItem('ecolink_session_token');
		}
	}
	
	isLoading.set(false);
}

export function login() {
	if (!browser || !window) return;
	
	const params = new URLSearchParams({
		client_id: GOOGLE_CLIENT_ID,
		redirect_uri: REDIRECT_URI,
		response_type: 'code',
		scope: 'openid profile email',
		access_type: 'offline',
		prompt: 'consent'
	});
	
	window.location.href = `https://accounts.google.com/o/oauth2/v2/auth?${params}`;
}

export async function handleCallback(code: string) {
	try {
		const response = await fetch(`${API_URL}/auth/google/callback`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify({
				code,
				redirect_uri: REDIRECT_URI
			})
		});

		if (!response.ok) {
			throw new Error('Falha na autenticação');
		}

		const data = await response.json();
		const userData = {
			id: data.user.id,
			name: data.user.name,
			email: data.user.email,
			picture: data.user.picture,
			provider: data.user.provider,
			created_at: data.user.created_at
		};

		// Save session token
		localStorage.setItem('ecolink_session_token', data.session_token);
		localStorage.setItem('ecolink_user', JSON.stringify(userData));

		user.set(userData);
		isAuthenticated.set(true);

		return userData;
	} catch (error) {
		console.error('Erro no callback:', error);
		throw error;
	}
}

export function logout() {
	isAuthenticated.set(false);
	user.set(null);
	if (browser) {
		localStorage.removeItem('ecolink_user');
		localStorage.removeItem('ecolink_session_token');
		// Redirect to home after logout
		window.location.href = '/';
	}
}

export function getSessionToken() {
	if (browser) {
		return localStorage.getItem('ecolink_session_token');
	}
	return null;
}