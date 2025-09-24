<script lang="ts">
	import { user, isAuthenticated } from '$lib/stores/auth';
	import { goto } from '$app/navigation';
	import { config } from '$lib/config/env';
	import { validateUserInput } from '$lib/utils/validation';
	
	let url = '';
	let loading = false;

	async function shortenUrl() {
		const validation = validateUserInput(url);
		if (!validation.isValid) {
			console.error('Invalid URL provided');
			return;
		}
		
		if (!$isAuthenticated || !$user) {
			const { login } = await import('$lib/services/authService');
			const state = btoa(JSON.stringify({ url: validation.sanitized }));
			login(state);
			return;
		}
		
		loading = true;
		try {
			const response = await fetch(`${config.apiUrl}/api/v1/links`, {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				credentials: 'include',
				body: JSON.stringify({ url })
			});
			
			if (response.ok) {
				const result = await response.json();
				const { pendingLink } = await import('$lib/stores/auth');
				pendingLink.set(result);
				goto('/result');
			} else {
				console.error('API Error:', response.status);
			}
		} catch (error) {
			console.error('Error shortening URL:', error);
		} finally {
			loading = false;
		}
	}
</script>

<svelte:head>
	<title>EcoLink - Sustainable Shortener</title>
</svelte:head>

<div class="text-center">
	<div class="max-w-3xl mx-auto">
		<h1 class="text-4xl font-bold text-gray-900 mb-4">
			Link shortening <span class="text-eco-400">Sustentável</span>
		</h1>
		<p class="text-xl text-gray-900 mb-8">
			Turn long URLs into short, eco-friendly links with included QR Code
		</p>
		
		<div class="bg-white rounded-xl shadow-lg p-8 mb-8">
			<form on:submit|preventDefault={shortenUrl} class="space-y-4">
				<div>
					<input
						bind:value={url}
						type="url"
						placeholder="Paste your URL here..."
						class="input-eco w-full text-lg"
						required
					/>
				</div>
				<button
					type="submit"
					disabled={loading}
					class="btn-eco text-lg px-8 py-3 disabled:opacity-50"
				>
					{loading ? 'Shortening...' : '🌱 Shorten Link'}
				</button>
			</form>
			

		</div>
		
		<div class="grid md:grid-cols-3 gap-8 text-center">
			<div class="bg-white p-6 rounded-lg shadow">
				<div class="text-3xl mb-2">⚡</div>
				<h3 class="font-semibold mb-2">Ultra Fast</h3>
				<p class="text-gray-600">Redirection in milliseconds</p>
			</div>
			<div class="bg-white p-6 rounded-lg shadow">
				<div class="text-3xl mb-2">🌱</div>
				<h3 class="font-semibold mb-2">Sustainable</h3>
				<p class="text-gray-600">Eco-friendly Infraestructure</p>
			</div>
			<div class="bg-white p-6 rounded-lg shadow">
				<div class="text-3xl mb-2">📊</div>
				<h3 class="font-semibold mb-2">Analytics</h3>
				<p class="text-gray-600">Detailed metrics</p>
			</div>
		</div>
	</div>
</div>