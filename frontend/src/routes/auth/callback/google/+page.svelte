<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { handleCallback } from '$lib/auth/google-direct';
	import type { User } from '$lib/stores/auth';


	let loading = true;
	let error: string | null = null;

	async function processPendingLink(userData: User) {
		const pendingUrl = localStorage.getItem('ecolink_pending_url');
		if (!pendingUrl) {
			goto('/dashboard');
			return;
		}

		try {
			const response = await fetch(`${import.meta.env.VITE_API_URL}/api/v1/links`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
					'X-User-ID': userData.id
				},
				body: JSON.stringify({ url: pendingUrl })
			});

			if (response.ok) {
				const result = await response.json();
				localStorage.removeItem('ecolink_pending_url');
				goto(`/result?short=${encodeURIComponent(result.shortUrl)}`);
			} else {
				localStorage.removeItem('ecolink_pending_url');
				goto('/dashboard');
			}
		} catch (err) {
			console.error('Error processing pending link:', err);
			localStorage.removeItem('ecolink_pending_url');
			goto('/dashboard');
		}
	}

	onMount(async () => {
		const code = $page.url.searchParams.get('code');
		
		if (!code) {
			error = 'Authorization code not found';
			loading = false;
			return;
		}

		try {
			// Processa o código OAuth real
			const userData = await handleCallback(code);
			
			// Check if there is a pending link to process
			await processPendingLink(userData);
			
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
			<p class="text-gray-600">Please wait a moment</p>
		{:else if error}
			<div class="text-red-600 mb-4">❌</div>
			<h1 class="text-xl font-semibold text-red-600 mb-2">Login Error</h1>
			<p class="text-gray-600 mb-4">{error}</p>
			<a href="/auth" class="btn-eco">Try Agaoin</a>
		{/if}
	</div>
</div>