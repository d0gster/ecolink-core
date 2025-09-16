import { writable } from 'svelte/store';
import { browser } from '$app/environment';
import type { Link } from '../types/link';

export interface User {
	id: string;
	name: string;
	email: string;
	picture: string;
	provider: string;
}

// Store para dados do usuário
export const user = writable<User | null>(null);

// Store para link pendente (antes do login)
export const pendingLink = writable<Link | null>(null);

// Funções de autenticação
export const auth = {
	// Login com Google OAuth real
	async loginWithProvider(provider: string) {
		if (provider.toLowerCase() === 'google') {
			const { getGoogleAuthUrl } = await import('$lib/auth/google.ts');
			window.location.href = getGoogleAuthUrl();
		} else {
			// Mock para outros provedores
			const mockUser = {
				id: `${provider}_${Date.now()}`,
				name: `User ${provider}`,
				email: `user@${provider}.com`,
				picture: `https://ui-avatars.com/api/?name=User+${provider}&background=16a34a&color=fff`,
				provider: provider.toLowerCase()
			};
			user.set(mockUser);
			if (browser) {
				localStorage.setItem('ecolink_user', JSON.stringify(mockUser));
			}
			return mockUser;
		}
	},

	// Logout
	logout() {
		user.set(null);
		pendingLink.set(null);
		if (browser) {
			localStorage.removeItem('ecolink_user');
			// Redirect to home after logout
			window.location.href = '/';
		}
	},

	// Save user after OAuth
	setUser(userData: User) {
		user.set(userData);
		if (browser) {
			localStorage.setItem('ecolink_user', JSON.stringify(userData));
		}
	},

	// Check if logged
	checkAuth() {
		if (browser) {
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
	}
};