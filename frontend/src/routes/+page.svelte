<script>
	import { user, isAuthenticated } from '$lib/auth/google-direct.js';

	import { goto } from '$app/navigation';
	
	let url = '';
	let loading = false;

	async function shortenUrl() {
		if (!url) return;
		
		// Se não está logado, salva o link e redireciona para login
		if (!$isAuthenticated || !$user) {
			localStorage.setItem('ecolink_pending_url', url);
			goto('/auth');
			return;
		}
		
		// Se está logado, cria o link diretamente
		loading = true;
		try {
			const response = await fetch('http://localhost:8080/api/v1/links', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
					'X-User-ID': $user.id
				},
				body: JSON.stringify({ url })
			});
			
			if (response.ok) {
				const result = await response.json();
				// Salva resultado completo no pendingLink
				const { pendingLink } = await import('$lib/stores/auth.js');
				pendingLink.set(result); // Salva resultado completo
				goto('/result');
			} else {
				console.error('Erro na API:', response.status);
			}
		} catch (error) {
			console.error('Erro ao encurtar URL:', error);
		} finally {
			loading = false;
		}
	}
</script>

<svelte:head>
	<title>EcoLink - Encurtador Sustentável</title>
</svelte:head>

<div class="text-center">
	<div class="max-w-3xl mx-auto">
		<h1 class="text-4xl font-bold text-gray-900 mb-4">
			Encurtamento de Links <span class="text-eco-400">Sustentável</span>
		</h1>
		<p class="text-xl text-gray-900 mb-8">
			Transforme URLs longas em links curtos e ecológicos com QR Code incluído
		</p>
		
		<div class="bg-white rounded-xl shadow-lg p-8 mb-8">
			<form on:submit|preventDefault={shortenUrl} class="space-y-4">
				<div>
					<input
						bind:value={url}
						type="url"
						placeholder="Cole sua URL aqui..."
						class="input-eco w-full text-lg"
						required
					/>
				</div>
				<button
					type="submit"
					disabled={loading}
					class="btn-eco text-lg px-8 py-3 disabled:opacity-50"
				>
					{loading ? 'Encurtando...' : '🌱 Encurtar Link'}
				</button>
			</form>
			

		</div>
		
		<div class="grid md:grid-cols-3 gap-8 text-center">
			<div class="bg-white p-6 rounded-lg shadow">
				<div class="text-3xl mb-2">⚡</div>
				<h3 class="font-semibold mb-2">Ultra Rápido</h3>
				<p class="text-gray-600">Redirecionamento em milissegundos</p>
			</div>
			<div class="bg-white p-6 rounded-lg shadow">
				<div class="text-3xl mb-2">🌱</div>
				<h3 class="font-semibold mb-2">Sustentável</h3>
				<p class="text-gray-600">Infraestrutura eco-friendly</p>
			</div>
			<div class="bg-white p-6 rounded-lg shadow">
				<div class="text-3xl mb-2">📊</div>
				<h3 class="font-semibold mb-2">Analytics</h3>
				<p class="text-gray-600">Métricas detalhadas</p>
			</div>
		</div>
	</div>
</div>