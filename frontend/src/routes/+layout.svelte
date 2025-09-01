<script>
	import '../app.css';
	import BackgroundVideo from '$lib/BackgroundVideo.svelte';
	import UserDropdown from '$lib/UserDropdown.svelte';
	import { initGoogleAuth } from '$lib/auth/google-direct.js';
	import { onMount } from 'svelte';

	onMount(() => {
		initGoogleAuth();
	});
</script>

<div class="min-h-screen relative">
	<!-- Vídeo de fundo -->
	<BackgroundVideo 
		desktopSrc="/videos/eco-background-desktop.mp4"
		mobileSrc="/videos/eco-background-mobile.mp4"
		poster=""
	/>
	<nav class="bg-white shadow-sm border-b border-eco-200">
		<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
			<div class="flex justify-between h-16">
				<div class="flex items-center">
					<h1 class="text-2xl font-bold text-eco-700">🌱 EcoLink</h1>
				</div>
				<div class="flex items-center space-x-4">
					<a href="/" class="text-gray-700 hover:text-eco-600">Home</a>
					<a href="/dashboard" class="text-gray-700 hover:text-eco-600">Dashboard</a>
					<UserDropdown />
				</div>
			</div>
		</div>
	</nav>
	
	<main class="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
		<slot />
	</main>
</div>