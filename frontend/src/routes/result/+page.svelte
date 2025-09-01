<script>
	import { onMount } from 'svelte';
	import { user, isAuthenticated } from '$lib/auth/google-direct.js';
	import { pendingLink } from '$lib/stores/auth.js';
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';

	let result = null;
	let loading = true;

	onMount(async () => {
		if (!$isAuthenticated || !$user) {
			goto('/auth');
			return;
		}

		// Pega resultado do pendingLink ou URL
		const pending = $pendingLink;
		if (pending) {
			result = pending;
			pendingLink.set(null); // Limpa após usar
			loading = false;
		} else {
			const shortUrl = $page.url.searchParams.get('short');
			if (shortUrl) {
				result = { short_url: shortUrl };
				loading = false;
			} else {
				goto('/dashboard');
			}
		}
	});
</script>

<svelte:head>
	<title>Link Criado - EcoLink</title>
</svelte:head>

<div class="text-center">
	<div class="max-w-3xl mx-auto">
		{#if loading}
			<div class="py-12">
				<div class="inline-block animate-spin rounded-full h-12 w-12 border-b-2 border-eco-600"></div>
				<p class="mt-4 text-gray-600">Criando seu link sustentável...</p>
			</div>
		{:else if result}
			<h1 class="text-3xl font-bold text-gray-900 mb-4">
				Link criado com <span class="text-eco-600">sucesso!</span>
			</h1>
			
			<div class="bg-white rounded-xl shadow-lg p-8 mb-8">
				<div class="space-y-6">
					<div>
						<label class="block text-sm font-medium text-gray-700 mb-2">URL Encurtada:</label>
						<div class="flex items-center space-x-2">
							<input
								value={result.short_url}
								readonly
								class="input-eco flex-1 text-lg"
							/>
							<button
								on:click={() => navigator.clipboard.writeText(result.short_url)}
								class="btn-eco"
							>
								Copiar
							</button>
						</div>
					</div>
					
					<div class="text-center">
						<p class="text-sm text-gray-600 mb-4">QR Code:</p>
						{#if result.qr_code}
							<img
								src="data:image/png;base64,{result.qr_code}"
								alt="QR Code"
								class="mx-auto border rounded-lg shadow w-48 h-48"
							/>
						{:else}
							<div class="w-48 h-48 mx-auto bg-gray-100 border rounded-lg flex items-center justify-center">
								<span class="text-gray-500">QR Code não disponível</span>
							</div>
						{/if}
					</div>
				</div>
			</div>

			<div class="flex justify-center space-x-4">
				<a href="/" class="btn-eco">Criar Outro Link</a>
				<a href="/dashboard" class="bg-gray-500 hover:bg-gray-600 text-white font-medium py-2 px-4 rounded-lg">
					Ver Dashboard
				</a>
			</div>
		{:else}
			<div class="py-12">
				<p class="text-red-600">Erro ao criar o link. Tente novamente.</p>
				<a href="/" class="btn-eco mt-4 inline-block">Voltar</a>
			</div>
		{/if}
	</div>
</div>