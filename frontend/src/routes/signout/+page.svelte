<script lang="ts">
	import { goto, invalidateAll } from '$app/navigation';
	import { isAuthenticated } from '$lib/stores/user';
	import { onMount } from 'svelte';

	let countdown = 1; // Countdown in seconds before redirection

	onMount(() => {
		// Start a countdown timer
		const interval = setInterval(async () => {
			countdown -= 1;
			if (countdown === 0) {
				clearInterval(interval);
				await handleSignOut(); // Sign out and redirect when countdown reaches 0
			}
		}, 1000);
	});

	// Handle sign out (you would also want to call your backend API to log out)
	async function handleSignOut() {
		await invalidateAll();
		isAuthenticated.set(false);
		await goto('/', { invalidateAll: true });
		// Redirect to the home page (or login page)
	}
</script>

<!-- Sign Out Page Layout -->
<div class="min-h-screen flex items-center justify-center bg-gray-100">
	<div class="text-center bg-white p-8 shadow-lg rounded-md">
		<h1 class="text-2xl font-bold mb-4">Signing out...</h1>
		<p class="text-gray-700 mb-4">
			You are being signed out and will be redirected in {countdown} seconds.
		</p>
		<p class="text-sm text-gray-500">Thank you for using our service!</p>
	</div>
</div>
