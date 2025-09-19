import { writable } from 'svelte/store';
import type { Link } from '../types/link';

export interface User {
	id: string;
	name: string;
	email: string;
	picture: string;
}

export const user = writable<User | null>(null);
export const isAuthenticated = writable(false);
export const isLoading = writable(true);
export const pendingLink = writable<Link | null>(null);