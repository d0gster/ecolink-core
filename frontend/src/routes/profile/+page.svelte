<script>
	import { onMount } from 'svelte';
	import { user, isAuthenticated, logout } from '$lib/auth/google-direct.js';
	import { goto } from '$app/navigation';

	let profileData = null;
	let loading = true;
	let error = null;

	onMount(async () => {
		if (!$isAuthenticated || !$user) {
			goto('/auth');
			return;
		}

		try {
			const response = await fetch('http://localhost:8080/api/v1/profile', {
				headers: {
					'X-User-ID': $user.id
				}
			});

			if (response.ok) {
				const data = await response.json();
				profileData = data.user;
			} else if (response.status === 404) {
				error = 'Usuário não encontrado. Faça login novamente.';
			} else {
				error = 'Erro ao carregar perfil';
			}
		} catch (err) {
			console.error('Erro:', err);
			error = 'Erro ao carregar perfil';
		} finally {
			loading = false;
		}
	});

	function handleLogout() {
		logout();
		goto('/');
	}
</script>

<svelte:head>
	<title>Perfil - EcoLink</title>
</svelte:head>

{#if loading}
	<div class="text-center py-8">
		<div class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-eco-600"></div>
		<p class="mt-2 text-gray-600">Carregando perfil...</p>
	</div>
{:else if error}
	<div class="text-center py-8">
		<p class="text-red-600">{error}</p>
	</div>
{:else if profileData}
	<div class="max-w-2xl mx-auto">
		<div class="bg-white rounded-xl shadow-lg overflow-hidden">
			<div class="bg-gradient-to-r from-eco-500 to-green-600 px-6 py-8">
				<div class="flex items-center space-x-4">
					<img
						src={profileData.picture || 'https://ui-avatars.com/api/?name=' + encodeURIComponent(profileData.name || profileData.email || 'Usuario') + '&background=16a34a&color=fff&size=80'}
						alt={profileData.name || 'Usuario'}
						class="w-20 h-20 rounded-full border-4 border-white shadow-lg"
						onerror="this.src='https://ui-avatars.com/api/?name=Usuario&background=16a34a&color=fff&size=80'"
					/>
					<div class="text-white">
						<h1 class="text-2xl font-bold">{profileData.name}</h1>
						<p class="text-green-100">EcoMembro desde {new Date(profileData.created_at).toLocaleDateString('pt-BR')}</p>
						<div class="flex items-center space-x-2 mt-2">
							<span class="px-2 py-1 bg-white/20 rounded text-sm capitalize">
								{profileData.provider}
							</span>
						</div>
					</div>
				</div>
			</div>

			<div class="p-6">
				<h2 class="text-xl font-semibold text-gray-900 mb-6">Informações do Google</h2>

				<div class="space-y-4">
					<div>
						<label class="block text-sm font-medium text-gray-700 mb-1">Nome</label>
						<p class="text-gray-900 py-2">{profileData.name}</p>
					</div>

					<div>
						<label class="block text-sm font-medium text-gray-700 mb-1">Email</label>
						<p class="text-gray-900 py-2">{profileData.email}</p>
					</div>

					<div>
						<label class="block text-sm font-medium text-gray-700 mb-1">ID do Google</label>
						<p class="text-gray-600 py-2 font-mono text-sm">{profileData.google_id}</p>
					</div>

					<div>
						<label class="block text-sm font-medium text-gray-700 mb-1">ID no Sistema</label>
						<p class="text-gray-600 py-2 font-mono text-sm">{profileData.id}</p>
					</div>

					<div>
						<label class="block text-sm font-medium text-gray-700 mb-1">Provedor</label>
						<p class="text-gray-900 py-2 capitalize">{profileData.provider}</p>
					</div>

					<div>
						<label class="block text-sm font-medium text-gray-700 mb-1">Última Atualização</label>
						<p class="text-gray-600 py-2">{new Date(profileData.updated_at).toLocaleString('pt-BR')}</p>
					</div>
				</div>

				<div class="mt-8 pt-6 border-t border-gray-200">
					<h3 class="text-lg font-semibold text-gray-900 mb-4">Ações da Conta</h3>
					<div class="space-y-3">
						<a href="/dashboard" class="block w-full text-center bg-gray-100 hover:bg-gray-200 text-gray-700 py-2 px-4 rounded-lg transition-colors">
							Ver Dashboard
						</a>
						<button
							on:click={handleLogout}
							class="w-full bg-red-500 hover:bg-red-600 text-white py-2 px-4 rounded-lg transition-colors"
						>
							Sair da Conta
						</button>
					</div>
				</div>
			</div>
		</div>
	</div>
{/if}