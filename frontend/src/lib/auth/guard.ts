import { browser } from '$app/environment';
import { goto } from '$app/navigation';
import { user } from '$lib/stores/auth.ts';
import type { User } from '$lib/stores/auth.ts';
import { get } from 'svelte/store';

// Check authenticated user
export function isAuthenticated() {
	if (!browser) return false;
	
	const currentUser: User | null = get(user);
	if (currentUser) return true;
	
	// Check localStorage
	const stored = localStorage.getItem('ecolink_user');
	if (stored) {
		try {
			const userData: User = JSON.parse(stored);
			user.set(userData);
			return true;
		} catch (e) {
			console.error("Error parsing user data from localStorage:", e);
			localStorage.removeItem('ecolink_user');
		}
	}
	
	return false;
}

// Middleware de proteção de rotas
export function requireAuth() {
	if (!isAuthenticated()) {
		goto('/auth');
		return false;
	}
	return true;
}

// Redirect logged user to dashboard
export function redirectIfAuthenticated() {
	if (isAuthenticated()) {
		goto('/dashboard');
		return true;
	}
	return false;
}