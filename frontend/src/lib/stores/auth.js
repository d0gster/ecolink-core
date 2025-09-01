import { writable } from 'svelte/store';
import { browser } from '$app/environment';

// Store para dados do usuário
export const user = writable(null);

// Store para link pendente (antes do login)
export const pendingLink = writable(null);

// Funções de autenticação
export const auth = {
	// Login com Google OAuth real
	async loginWithProvider(provider) {
		if (provider.toLowerCase() === 'google') {
			const { getGoogleAuthUrl } = await import('$lib/auth/google.js');
			window.location.href = getGoogleAuthUrl();
		} else {
			// Mock para outros provedores
			const mockUser = {
				id: `${provider}_${Date.now()}`,
				name: `User ${provider}`,
				email: `user@${provider}.com`,
				image: `https://ui-avatars.com/api/?name=User+${provider}&background=16a34a&color=fff`,
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
			// Redireciona para home após logout
			window.location.href = '/';
		}
	},

	// Salva usuário após OAuth
	setUser(userData) {
		user.set(userData);
		if (browser) {
			localStorage.setItem('ecolink_user', JSON.stringify(userData));
		}
	},

	// Verifica se está logado
	checkAuth() {
		if (browser) {
			const stored = localStorage.getItem('ecolink_user');
			if (stored) {
				user.set(JSON.parse(stored));
				return true;
			}
		}
		return false;
	}
};