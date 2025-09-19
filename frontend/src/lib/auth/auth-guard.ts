import { goto } from '$app/navigation';
import { user, isAuthenticated, isLoading } from '$lib/stores/auth';
import { get } from 'svelte/store';

export async function requireAuth(): Promise<boolean> {
	if (get(isLoading)) {
		return new Promise((resolve) => {
			const unsubscribe = isLoading.subscribe((loading) => {
				if (!loading) {
					unsubscribe();
					resolve(checkAuth());
				}
			});
		});
	}
	
	return checkAuth();
}

function checkAuth(): boolean {
	if (!get(isAuthenticated) || !get(user)) {
		goto('/auth');
		return false;
	}
	return true;
}