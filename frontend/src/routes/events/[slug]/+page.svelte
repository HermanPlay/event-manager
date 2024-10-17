<script lang="ts">
	import { PUBLIC_API_URL } from '$env/static/public';
	import type { Event } from '$lib/schemas/event';
	import { addToast } from '$lib/stores/toasts.js';
	import SvelteMarkdown from 'svelte-markdown';

	export let data;

	let event: Event = data.event ?? { title: '', description: '', location: '', date: '' };
	// Get the event ID from the URL

	// Function to handle booking the event
	async function bookEvent() {
		try {
			const response = await fetch(`${PUBLIC_API_URL}/event/book/${event.id}`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
					Authorization: 'Bearer ' + data.token
				}
			});

			if (!response.ok) {
				if (response.status === 409) {
					addToast({
						type: 'info',
						message: "You've already booked this event.",
						timeout: 5000
					});
					return;
				} else {
					throw new Error('Failed to book event.');
				}
			}

			addToast({
				type: 'success',
				message: 'Event booked successfully.',
				timeout: 2000
			});
		} catch (error: any) {
			addToast({
				type: 'error',
				message: error.message,
				timeout: 5000
			});
		}
	}
	console.log(event.description);
</script>

<!-- Page Layout -->

<div class="min-h-screen p-8 bg-gray-100">
	{#if data.error}
		<p class="text-red-500">{data.error}</p>
	{:else}
		<div class="container mx-auto p-6 bg-white shadow-lg rounded-lg">
			<h1 class="text-3xl font-bold mb-2">{event.title}</h1>
			<p class="text-l text-gray-800 mb-4">{event.short_description}</p>

			<div class="mb-4">
				<p class="text-gray-700 text-sm">Date: {event.date}</p>
				<p class="text-gray-700 text-sm">Location: {event.location}</p>
			</div>

			<div class="mb-6">
				<h2 class="text-2xl font-semibold mb-4">Event Details</h2>
				<SvelteMarkdown source={event.description} />
			</div>

			<div class="flex justify-end">
				<button
					class="bg-blue-600 hover:bg-blue-700 text-white py-2 px-6 rounded"
					on:click={bookEvent}
				>
					Book Now
				</button>
			</div>
		</div>
	{/if}
</div>

<style>
	/* Basic styling for the event page */
	.container {
		max-width: 800px;
		margin: 0 auto;
	}
</style>
