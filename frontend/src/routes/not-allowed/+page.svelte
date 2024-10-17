<script lang="ts">
	import { goto } from '$app/navigation';

	// Function to redirect the user to the homepage
	function goToHome() {
		goto('/');
	}
</script>

<div class="min-h-screen flex items-center justify-center bg-gray-100">
	<div class="bg-white shadow-md rounded-lg p-8 max-w-lg text-center">
		<h1 class="text-4xl font-bold text-red-600 mb-4">Access Denied</h1>
		<p class="text-gray-700 mb-6">You don't have permission to access this page.</p>
		<button
			on:click={goToHome}
			class="bg-blue-600 hover:bg-blue-700 text-white py-2 px-4 rounded transition-colors"
		>
			Go to Homepage
		</button>
	</div>
</div>

<style>
	.min-h-screen {
		min-height: 100vh;
	}
</style>
