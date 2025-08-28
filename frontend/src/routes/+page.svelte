<script>
	let url = '';
	let result = null;
	let loading = false;

	async function shortenUrl() {
		if (!url) return;
		
		loading = true;
		try {
			const response = await fetch('http://localhost:8080/api/v1/links', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
					'X-User-ID': 'demo-user'
				},
				body: JSON.stringify({ url })
			});
			
			if (response.ok) {
				result = await response.json();
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
			Encurtamento de Links <span class="text-eco-600">Sustentável</span>
		</h1>
		<p class="text-xl text-gray-600 mb-8">
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
			
			{#if result}
				<div class="mt-8 p-6 bg-eco-50 rounded-lg border border-eco-200">
					<h3 class="text-lg font-semibold text-eco-800 mb-4">Link criado com sucesso!</h3>
					<div class="space-y-4">
						<div>
							<label class="block text-sm font-medium text-gray-700 mb-1">URL Encurtada:</label>
							<div class="flex items-center space-x-2">
								<input
									value={result.short_url}
									readonly
									class="input-eco flex-1"
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
							<p class="text-sm text-gray-600 mb-2">QR Code:</p>
							<img
								src="data:image/png;base64,{result.qr_code}"
								alt="QR Code"
								class="mx-auto border rounded"
							/>
						</div>
					</div>
				</div>
			{/if}
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