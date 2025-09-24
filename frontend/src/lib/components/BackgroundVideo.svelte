<script lang="ts">
	interface Props {
		desktopSrc?: string;
		mobileSrc?: string;
		poster?: string;
	}
	
	let { desktopSrc = '', mobileSrc = '', poster = '' }: Props = $props();
	
	let videoElement: HTMLVideoElement;
	let isMobile = $state(false);
	
	// Svelte 5: usando $effect para detectar mobile
	$effect(() => {
		if (typeof window !== 'undefined') {
			isMobile = window.innerWidth < 768;
		}
	});
</script>

<div class="fixed inset-0 -z-10 overflow-hidden">
	<video
		bind:this={videoElement}
		class="absolute inset-0 w-full h-full object-cover opacity-70"
		autoplay
		muted
		loop
		playsinline
		{poster}
		onloadeddata={() => videoElement?.play()}
	>
		<source src={isMobile ? mobileSrc : desktopSrc} type="video/mp4" />
		<!-- Fallback para navegadores sem suporte -->
		<div class="absolute inset-0 bg-gradient-to-br from-green-50 to-green-100"></div>
	</video>
	
	<!-- Overlay para melhor legibilidade -->
	<div class="absolute inset-0 bg-white/20 backdrop-blur-[0.5px]"></div>
</div>

<style>
	video {
		filter: saturate(1.2) brightness(0.9);
	}
</style>