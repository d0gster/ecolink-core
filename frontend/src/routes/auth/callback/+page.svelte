<script>
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { handleCallback } from '$lib/auth/google-direct.js';
	import { page } from '$app/stores';

	let loading = true;
	let error = null;

	onMount(async () => {
		const code = $page.url.searchParams.get('code');
		
		if (!code) {
			error = 'Código de autorização não encontrado';
			loading = false;
			return;
		}

		try {
			await handleCallback(code);
			
			// Verifica se há URL pendente
			const pendingUrl = localStorage.getItem('ecolink_pending_url');
			if (pendingUrl) {
				localStorage.removeItem('ecolink_pending_url');
				goto(`/?url=${encodeURIComponent(pendingUrl)}`);
			} else {
				goto('/dashboard');
			}
		} catch (err) {
			console.error('Erro no callback:', err);
			error = err.message;
			loading = false;
		}
	});
</script>

<svelte:head>
	<title>Processando Login - EcoLink</title>
</svelte:head>

<div class="min-h-[80vh] flex items-center justify-center">
	<div class="text-center">
		{#if loading}
			<div class="inline-block animate-spin rounded-full h-12 w-12 border-b-2 border-eco-600 mb-4"></div>
			<h1 class="text-xl font-semibold text-gray-900 mb-2">Processando login...</h1>
			<p class="text-gray-600">Aguarde um momento</p>
		{:else if error}
			<div class="text-red-600 mb-4">❌</div>
			<h1 class="text-xl font-semibold text-red-600 mb-2">Erro no Login</h1>
			<p class="text-gray-600 mb-4">{error}</p>
			<a href="/auth" class="btn-eco">Tentar Novamente</a>
		{/if}
	</div>
</div>