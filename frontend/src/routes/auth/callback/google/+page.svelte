<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { handleCallback } from '$lib/services/authService';

	let loading = true;
	let error: string | null = null;

	onMount(async () => {
		const code = $page.url.searchParams.get('code');
		const state = $page.url.searchParams.get('state');

		if (!code) {
			error = 'Authorization code not found';
			loading = false;
			return;
		}

		// Verify state parameter for CSRF protection
		const storedState = sessionStorage.getItem('oauth_state');
		if (!state || state !== storedState) {
			error = 'Invalid state parameter - possible CSRF attack';
			loading = false;
			return;
		}

		// Clear stored state
		sessionStorage.removeItem('oauth_state');

		try {
			await handleCallback(code, state);

			if (state) {
				try {
					const { url: pendingUrl } = JSON.parse(atob(state));
					if (pendingUrl) {
						const response = await fetch(`${import.meta.env.VITE_API_URL}/api/v1/links`, {
							method: 'POST',
							headers: { 'Content-Type': 'application/json' },
							credentials: 'include',
							body: JSON.stringify({ url: pendingUrl })
						});

						if (response.ok) {
							const result = await response.json();
							const { pendingLink } = await import('$lib/stores/auth');
							pendingLink.set(result);
							goto('/result');
							return;
						}
					}
				} catch (err) {
					console.error('Error processing state:', err);
				}
			}
			
			goto('/dashboard');
		} catch (err: any) {
			console.error('Callback error:', err);
			error = err.message || 'Authentication error';
			loading = false;
		}
	});
</script>

<svelte:head>
	<title>Processing Login - EcoLink</title>
</svelte:head>

<div class="min-h-[80vh] flex items-center justify-center">
	<div class="text-center">
		{#if loading}
			<div class="inline-block animate-spin rounded-full h-12 w-12 border-b-2 border-eco-600 mb-4"></div>
			<h1 class="text-xl font-semibold text-gray-900 mb-2">Processing login...</h1>
			<p class="text-gray-600">Wait a moment</p>
		{:else if error}
			<div class="text-red-600 mb-4">❌</div>
			<h1 class="text-xl font-semibold text-red-600 mb-2">Login error</h1>
			<p class="text-gray-600 mb-4">{error}</p>
			<a href="/auth" class="btn-eco">Try again</a>
		{/if}
	</div>
</div>