<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { handleCallback } from '$lib/auth/google-direct.ts';

	let loading = true;
	let error: string | null = null;

	onMount(async () => {
		const code = $page.url.searchParams.get('code');
		
		if (!code) {
			error = 'Authorization code not found';
			loading = false;
			return;
		}

		try {
			await handleCallback(code);
			
			// Check for pending URL
			const pendingUrl = localStorage.getItem('ecolink_pending_url');
			if (pendingUrl) {
				localStorage.removeItem('ecolink_pending_url');
				goto(`/?url=${encodeURIComponent(pendingUrl)}`);
			} else {
				goto('/dashboard');
			}
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