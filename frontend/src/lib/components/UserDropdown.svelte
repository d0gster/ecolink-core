<script lang="ts">
	import { getContext } from 'svelte';
	import { user, isAuthenticated, logout } from '$lib/auth/google-direct.ts';
	import { goto } from '$app/navigation';
	
	let showDropdown = false;

	function handleLogout() {
		logout();
		showDropdown = false;
	}

	function goToProfile() {
		showDropdown = false;
		goto('/profile');
	}

	function handleClickOutside(event: MouseEvent) {
		if (event.target instanceof Element && !event.target.closest('.relative')) {
			showDropdown = false;
		}
	}
</script>

{#if $isAuthenticated && $user}
	<div class="relative">
		<button
			on:click|stopPropagation={() => showDropdown = !showDropdown}
			class="flex items-center space-x-2 text-gray-700 hover:text-eco-600 focus:outline-none"
		>
			<img
				src={$user.picture || 'https://ui-avatars.com/api/?name=' + encodeURIComponent($user.name || $user.email || 'Usuario') + '&background=16a34a&color=fff'}
				alt={$user.name || 'Usuario'}
				class="w-8 h-8 rounded-full"
				referrerpolicy="no-referrer"
				crossorigin="anonymous"
			/>
			<span class="hidden md:block font-medium">{$user.name}</span>
			<svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
				<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
			</svg>
		</button>

		{#if showDropdown}
			<div class="absolute right-0 mt-2 w-48 bg-white rounded-lg shadow-lg border border-gray-200 z-50" on:click|stopPropagation>
				<div class="py-2">
					<div class="px-4 py-2 border-b border-gray-100">
						<p class="text-sm font-medium text-gray-900">{$user.name}</p>
						<p class="text-xs text-gray-500">{$user.email}</p>
					</div>
					
					<button
						on:click={goToProfile}
						class="w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-50 flex items-center space-x-2"
					>
						<span>👤</span>
						<span>Perfil</span>
					</button>
					
					<button
						on:click={handleLogout}
						class="w-full text-left px-4 py-2 text-sm text-red-600 hover:bg-red-50 flex items-center space-x-2"
					>
						<span>🚪</span>
						<span>Sair</span>
					</button>
				</div>
			</div>
		{/if}
	</div>
{:else}
	<a href="/auth" class="text-gray-700 hover:text-eco-600 font-medium">
		Entrar
	</a>
{/if}

<!-- Fecha dropdown ao clicar fora -->
<svelte:window on:click={handleClickOutside} />