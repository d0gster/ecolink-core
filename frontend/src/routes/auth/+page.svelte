<script lang="ts">
	import { login } from '$lib/auth/google-direct';
	import { isAuthenticated } from '$lib/stores/auth';
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import type { Provider } from '$lib/types/provider';

	let loading = false;

	onMount(() => {
		// Redirects if already logged in
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

	async function loginWith(provider: Provider) {
		if (!provider.functional) {
			alert(`Login with ${provider.name} will be implemented soon!`);
			return;
		}
		
		loading = true;
		try {
			const pendingUrl = localStorage.getItem('ecolink_pending_url');
			let state;
			if (pendingUrl) {
				state = btoa(JSON.stringify({ url: pendingUrl }));
			}
			await login(state);
		} catch (error) {
			console.error('Login error:', error);
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
			<h1 class="text-2xl font-bold text-gray-900 mb-2">Enter EcoLink</h1>
			<p class="text-gray-600">Log in to access your sustainable links</p>
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
						{loading && provider.functional ? 'Redirecting...' : `Continue with ${provider.name}`}
						{!provider.functional ? ' (Soon)' : ''}
					</span>
				</button>
			{/each}
		</div>

		<div class="mt-8 text-center text-sm text-gray-500">
			<p>By continuing, you agree to our</p>
			<p>
				<a href="/terms" class="text-eco-600 hover:text-eco-700">Terms of Use</a> and
				<a href="/privacy" class="text-eco-600 hover:text-eco-700">Privacy Policy</a>
			</p>
		</div>
	</div>
</div>