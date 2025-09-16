<script lang="ts">
	import { onMount } from 'svelte';
	import { user, isAuthenticated } from '$lib/auth/google-direct';	
	import { pendingLink } from '$lib/stores/auth';
	import type { Result } from '$lib/types/result';
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	
	let result: Result = {shortUrl: ""};
	let loading = true;

	onMount(async () => {
		if (!$isAuthenticated || !$user) {
			goto('/auth');
			return;
		}

		// Get result from pendingLink or URL
		const pending = $pendingLink;
		if (pending) {
			result = { shortUrl: pending.shortUrl, qrCode: pending.qrCode };
			pendingLink.set(null); // Clear after use
			loading = false;
		} else {
			const shortUrl = $page.url.searchParams.get('short');
			if (shortUrl) {
				result = { shortUrl: shortUrl };
				loading = false;
			} else {
				goto('/dashboard');
			}
		}
	});
</script>

<svelte:head>
	<title>Link Created - EcoLink</title>
</svelte:head>

<div class="text-center">
	<div class="max-w-3xl mx-auto">
		{#if loading}
			<div class="py-12">
				<div class="inline-block animate-spin rounded-full h-12 w-12 border-b-2 border-eco-600"></div>
				<p class="mt-4 text-gray-600">Creating your sustainable link...</p>
			</div>
		{:else if result}
			<h1 class="text-3xl font-bold text-gray-900 mb-4">
				Link created <span class="text-eco-600">successfully!</span>
			</h1>
			
			<div class="bg-white rounded-xl shadow-lg p-8 mb-8">
				<div class="space-y-6">
					<div>
						<label for="shortened-url" class="block text-sm font-medium text-gray-700 mb-2">Shortened URL:</label>
						<div class="flex items-center space-x-2">
							{#if result}
							<input
								id="shortened-url"
									value={result.shortUrl}
									readonly
									class="input-eco flex-1 text-lg"
								/>
							<button
								on:click={() => navigator.clipboard.writeText(result.shortUrl)}
								class="btn-eco"
							>
								Copy
							</button>
							{/if}
						</div>
					</div>
					
					<div class="text-center">
						<p class="text-sm text-gray-600 mb-4">QR Code:</p>
						{#if result.qrCode}
							<img
								src="data:image/png;base64,{result.qrCode}"
								alt="QR Code"
								class="mx-auto border rounded-lg shadow w-48 h-48"
							/>
						{:else}
							<div class="w-48 h-48 mx-auto bg-gray-100 border rounded-lg flex items-center justify-center">
								<span class="text-gray-500">QR Code not available</span>
							</div>
						{/if}
					</div>
				</div>
			</div>

			<div class="flex justify-center space-x-4">
				<a href="/" class="btn-eco">Create Another Link</a>
				<a href="/dashboard" class="bg-gray-500 hover:bg-gray-600 text-white font-medium py-2 px-4 rounded-lg">
					View Dashboard
				</a>
			</div>
		{:else}
			<div class="py-12">
				<p class="text-red-600">Error creating link. Please try again.</p>
				<a href="/" class="btn-eco mt-4 inline-block">Go Back</a>
			</div>
		{/if}
	</div>
</div>