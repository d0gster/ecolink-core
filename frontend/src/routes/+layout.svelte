<script lang="ts">
	import '../app.css';
import BackgroundVideo from '$lib/components/BackgroundVideo.svelte';
import UserDropdown from '$lib/components/UserDropdown.svelte';
import { initGoogleAuth } from '$lib/auth/google-direct';
import { isLoading } from '$lib/stores/auth';
import { onDestroy } from 'svelte';
import { derived } from 'svelte/store';
	import { onMount } from 'svelte';

onMount(() => {
    // Initialize authentication on app load. This delegates to authService.verifySession
    // which will set user / isAuthenticated / isLoading appropriately.
    initGoogleAuth();
});

// derived store to show when skeleton should be visible
const showSkeleton = derived(isLoading, ($isLoading) => $isLoading === true);

let skeletonUnsub = showSkeleton.subscribe(() => {});
onDestroy(() => skeletonUnsub());
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
	
	{#if $isLoading}
		<!-- Skeleton while session is being verified to avoid UI flashes -->
		<main class="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
			<div class="animate-pulse space-y-4">
				<div class="h-6 bg-gray-200 rounded w-1/3"></div>
				<div class="h-4 bg-gray-200 rounded w-full"></div>
				<div class="h-48 bg-gray-200 rounded w-full"></div>
			</div>
		</main>
	{:else}
		<main class="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
			<slot />
		</main>
	{/if}
</div>