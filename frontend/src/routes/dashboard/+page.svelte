<script>
	import { onMount } from 'svelte';
	import { user, isAuthenticated } from '$lib/auth/google-direct.js';
	import { goto } from '$app/navigation';
	
	let links = [];
	let loading = true;

	onMount(async () => {
		if (!$isAuthenticated || !$user) {
			goto('/auth');
			return;
		}

		try {
			const response = await fetch('http://localhost:8080/api/v1/links', {
				headers: {
					'X-User-ID': $user.id
				}
			});
			
			if (response.ok) {
				const data = await response.json();
				links = data.links || [];
			}
		} catch (error) {
			console.error('Erro ao carregar links:', error);
		} finally {
			loading = false;
		}
	});
</script>

<svelte:head>
	<title>Dashboard - EcoLink</title>
</svelte:head>

<div>
	<div class="mb-8">
		<h1 class="text-3xl font-bold text-gray-900">Dashboard</h1>
		<p class="text-gray-700 mt-2">Gerencie seus links encurtados</p>
	</div>

	{#if loading}
		<div class="text-center py-8">
			<div class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-eco-600"></div>
			<p class="mt-2 text-gray-600">Carregando seus links...</p>
		</div>
	{:else if links.length === 0}
		<div class="text-center py-12 bg-white rounded-lg shadow">
			<div class="text-6xl mb-4">🌱</div>
			<h3 class="text-xl font-semibold text-gray-900 mb-2">Nenhum link ainda</h3>
			<p class="text-gray-600 mb-4">Comece criando seu primeiro link sustentável!</p>
			<a href="/" class="btn-eco">Criar Primeiro Link</a>
		</div>
	{:else}
		<div class="bg-white shadow rounded-lg overflow-hidden">
			<div class="px-6 py-4 border-b border-gray-200">
				<h2 class="text-lg font-semibold text-gray-900">Seus Links ({links.length})</h2>
			</div>
			<div class="divide-y divide-gray-200">
				{#each links as link}
					<div class="px-6 py-4 hover:bg-gray-50">
						<div class="flex items-center justify-between">
							<div class="flex-1 min-w-0">
								<div class="flex items-center space-x-3">
									<div class="flex-shrink-0">
										<div class="w-10 h-10 bg-eco-100 rounded-full flex items-center justify-center">
											<span class="text-eco-600 font-semibold">🔗</span>
										</div>
									</div>
									<div class="flex-1 min-w-0">
										<p class="text-sm font-medium text-gray-900 truncate">
											{link.original_url}
										</p>
										<div class="flex items-center space-x-4 mt-1">
											<p class="text-sm text-eco-600">
												localhost:8080/{link.short_code}
											</p>
											<span class="text-sm text-gray-500">
												{link.click_count} cliques
											</span>
											<span class="text-sm text-gray-500">
												{new Date(link.created_at).toLocaleDateString('pt-BR')}
											</span>
										</div>
									</div>
								</div>
							</div>
							<div class="flex items-center space-x-2">
								<button
									on:click={() => navigator.clipboard.writeText(`http://localhost:8080/${link.short_code}`)}
									class="text-eco-600 hover:text-eco-700 text-sm font-medium"
								>
									Copiar
								</button>
							</div>
						</div>
					</div>
				{/each}
			</div>
		</div>
	{/if}
</div>