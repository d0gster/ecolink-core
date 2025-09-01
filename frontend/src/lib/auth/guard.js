import { browser } from '$app/environment';
import { goto } from '$app/navigation';
import { user } from '$lib/stores/auth.js';
import { get } from 'svelte/store';

// Verifica se usuário está autenticado
export function isAuthenticated() {
	if (!browser) return false;
	
	const currentUser = get(user);
	if (currentUser) return true;
	
	// Verifica localStorage
	const stored = localStorage.getItem('ecolink_user');
	if (stored) {
		try {
			const userData = JSON.parse(stored);
			user.set(userData);
			return true;
		} catch (e) {
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

// Redireciona usuário logado para dashboard
export function redirectIfAuthenticated() {
	if (isAuthenticated()) {
		goto('/dashboard');
		return true;
	}
	return false;
}