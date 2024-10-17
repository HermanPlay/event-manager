<script lang="ts">
	import { goto } from '$app/navigation';
	import FeaturedEventCard from './components/FeaturedEventCard.svelte';
	import { isAuthenticated } from '$lib/stores/user';
	import { PUBLIC_API_URL } from '$env/static/public';
	import { onMount } from 'svelte';
	export let data;

	const getFeaturedEvents = async () => {
		// Make an API call to /event/featured
		// and return the response
		const response = await fetch(`${PUBLIC_API_URL}/event/featured`);

		if (!response.ok) {
			throw new Error('Failed to fetch featured events.');
		}
		const data = await response.json();
		return data.data;
	};
	onMount(async () => {
		const token = data.token;
		if (!token) {
			isAuthenticated.set(false);
		}
	});
</script>

<!-- Home Page Content -->
<div class="min-h-screen bg-gray-100">
	<!-- Hero Section -->
	<div class="hero text-white p-10 text-center items-center flex flex-col justify-center">
		<h1 class="text-4xl font-bold mb-4">Welcome to Event Manager</h1>
		<p class="text-lg">Find and book exciting events around you!</p>

		{#if $isAuthenticated}
			<button
				class="mt-6 bg-green-500 text-white px-4 py-2 rounded"
				on:click={() => {
					goto('/events');
				}}
			>
				View Your Events
			</button>
		{:else}
			<button
				class="mt-6 bg-white text-blue-500 px-4 py-2 rounded"
				on:click={() => {
					goto('/auth');
				}}
			>
				Book Now!
			</button>
		{/if}
	</div>

	<!-- Featured Events Section -->
	<div class="p-10">
		<h2 class="text-2xl font-semibold mb-6">Featured Events</h2>
		<div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
			{#await getFeaturedEvents() then events}
				{#each events as event}
					<FeaturedEventCard {event} />
				{/each}
			{:catch error}
				<p class="text-red-500 text-center">{error.message}</p>
			{/await}
		</div>
	</div>
</div>

<style>
	.hero {
		background-image: url('/hero-bg.jpg');
		background-position: center;
		background-repeat: no-repeat;
		background-size: cover;
		min-height: 20rem;
	}
</style>
