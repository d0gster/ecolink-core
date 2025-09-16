import { createAuth0Client, Auth0Client } from '@auth0/auth0-spa-js';
import { writable } from 'svelte/store';
import { browser } from '$app/environment';
import type { User } from '$lib/stores/auth.ts';

export const isLoading = writable(true);
export const isAuthenticated = writable(false);
export const authToken = writable('');
export const user = writable<User | null>(null);

let auth0Client: Auth0Client | undefined;

export async function initAuth0() {
	if (!browser) return;

	auth0Client = await createAuth0Client({
		domain: import.meta.env.VITE_AUTH0_DOMAIN,
		clientId: import.meta.env.VITE_AUTH0_CLIENT_ID,
		authorizationParams: {
			redirect_uri: window.location.origin + '/auth/callback',
			audience: import.meta.env.VITE_AUTH0_AUDIENCE
		}
	});

	const isAuth = await auth0Client.isAuthenticated();
	isAuthenticated.set(isAuth);

	if (isAuth) {
		const userData: User | undefined = await auth0Client.getUser();
		const token: string = await auth0Client.getTokenSilently();
		
		if (userData) {
			user.set(userData);
		}
		authToken.set(token);
	}

	isLoading.set(false);
}

export async function login() {
	if (!auth0Client) return;
	
	await auth0Client.loginWithRedirect({
		authorizationParams: {
			redirect_uri: window.location.origin + '/auth/callback'
		}
	});
}

export async function logout() {
	if (!auth0Client) return;
	
	await auth0Client.logout({
		logoutParams: {
			returnTo: window.location.origin
		}
	});
}

export async function handleCallback() {
	if (!auth0Client) return;
	
	await auth0Client.handleRedirectCallback();
	
	const isAuth = await auth0Client.isAuthenticated();
	isAuthenticated.set(isAuth);
	
	if (isAuth) {
		const userData: User | undefined = await auth0Client.getUser();
		const token: string = await auth0Client.getTokenSilently();
		
		if (userData) {
			user.set(userData);
		}
		authToken.set(token);
	}
}