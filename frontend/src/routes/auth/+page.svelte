<script>
	import { login, isAuthenticated } from '$lib/auth/google-direct.js';
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';

	let loading = false;

	onMount(() => {
		// Redireciona se já estiver logado
		if ($isAuthenticated) {
			goto('/dashboard');
		}
	});

	const providers = [
		{ name: 'Google', icon: '🔍', color: 'bg-red-500 hover:bg-red-600', functional: true },
		{ name: 'Microsoft', icon: '🪟', color: 'bg-blue-500 hover:bg-blue-600', functional: false },
		{ name: 'LinkedIn', icon: '💼', color: 'bg-blue-700 hover:bg-blue-800', functional: false },
		{ name: 'GitHub', icon: '🐙', color: 'bg-gray-800 hover:bg-gray-900', functional: false }
	];

	async function loginWith(provider) {
		if (!provider.functional) {
			alert(`Login com ${provider.name} será implementado em breve!`);
			return;
		}
		
		loading = true;
		try {
			await login();
		} catch (error) {
			console.error('Erro no login:', error);
			loading = false;
		}
	}
</script>

<svelte:head>
	<title>Login - EcoLink</title>
</svelte:head>

<div class="min-h-[80vh] flex items-center justify-center">
	<div class="max-w-md w-full bg-white rounded-xl shadow-lg p-8">
		<div class="text-center mb-8">
			<div class="text-4xl mb-4">🌱</div>
			<h1 class="text-2xl font-bold text-gray-900 mb-2">Entre no EcoLink</h1>
			<p class="text-gray-600">Faça login para acessar seus links sustentáveis</p>
		</div>

		<div class="space-y-3">
			{#each providers as provider}
				<button
					on:click={() => loginWith(provider)}
					disabled={loading}
					class="w-full flex items-center justify-center space-x-3 py-3 px-4 rounded-lg text-white font-medium transition-colors disabled:opacity-50 {provider.color} {!provider.functional ? 'opacity-75' : ''}"
				>
					<span class="text-xl">{provider.icon}</span>
					<span>
						{loading && provider.functional ? 'Redirecionando...' : `Continuar com ${provider.name}`}
						{!provider.functional ? ' (Em breve)' : ''}
					</span>
				</button>
			{/each}
		</div>

		<div class="mt-8 text-center text-sm text-gray-500">
			<p>Ao continuar, você concorda com nossos</p>
			<p>
				<a href="#" class="text-eco-600 hover:text-eco-700">Termos de Uso</a> e 
				<a href="#" class="text-eco-600 hover:text-eco-700">Política de Privacidade</a>
			</p>
		</div>
	</div>
</div>