<script>
	import '../app.css';
</script>

<div class="min-h-screen bg-gradient-to-br from-eco-50 to-green-100">
	<nav class="bg-white shadow-sm border-b border-eco-200">
		<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
			<div class="flex justify-between h-16">
				<div class="flex items-center">
					<h1 class="text-2xl font-bold text-eco-700">🌱 EcoLink</h1>
				</div>
				<div class="flex items-center space-x-4">
					<a href="/" class="text-gray-700 hover:text-eco-600">Home</a>
					<a href="/dashboard" class="text-gray-700 hover:text-eco-600">Dashboard</a>
				</div>
			</div>
		</div>
	</nav>
	
	<main class="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
		<slot />
	</main>
</div>