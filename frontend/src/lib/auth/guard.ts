import { browser } from '$app/environment';
import { goto } from '$app/navigation';
import { user, isLoading } from '$lib/stores/auth';
import type { User } from '$lib/stores/auth';
import { get } from 'svelte/store';

// Check authenticated user
export function isAuthenticated() {
	if (!browser) return false;
	
	const currentUser: User | null = get(user);
	return !!currentUser;
}

// Middleware de proteção de rotas
export function requireAuth() {
    // If session is still verifying, don't redirect yet
    if (get(isLoading)) return false;

    if (!isAuthenticated()) {
        goto('/auth');
        return false;
    }

    return true;
}

// Redirect logged user to dashboard
export function redirectIfAuthenticated() {
    // Wait for session verification to complete before redirecting
    if (get(isLoading)) return false;

    if (isAuthenticated()) {
        goto('/dashboard');
        return true;
    }

    return false;
}