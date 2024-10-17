<script lang="ts">
	import { goto } from '$app/navigation';
	import { PUBLIC_API_URL } from '$env/static/public';
	import type { Event } from '$lib/schemas/event';
	import { addToast } from '$lib/stores/toasts.js';

	export let data;
	let event = data.event
		? data.event
		: ({ title: '', description: '', location: '', time: '', short_description: '' } as Event);

	// Function to update the event
	async function updateEvent() {
		try {
			const response = await fetch(`${PUBLIC_API_URL}/event/${event.id}`, {
				method: 'PATCH',
				headers: {
					'Content-Type': 'application/json',
					Authorization: 'Bearer ' + data.token
				},
				body: JSON.stringify(event)
			});

			if (!response.ok) {
				console.error(await response.json());
				throw new Error('Failed to update event.');
			}

			// Redirect the user back to the event details page after successful update
			goto(`/events/${event.id}`);
		} catch (error: any) {
			addToast({
				message: error.message,
				type: 'error'
			});
		}
	}
</script>

<!-- Page Layout -->

<div class="min-h-screen p-8 bg-gray-100">
	<div class="container mx-auto p-6 bg-white shadow-lg rounded-lg">
		<h1 class="text-2xl font-bold mb-6">Edit Event</h1>

		<form on:submit|preventDefault={updateEvent}>
			<!-- Event Title -->
			<div class="mb-4">
				<label for="title" class="block text-gray-700">Title:</label>
				<input
					id="title"
					type="text"
					bind:value={event.title}
					class="w-full p-3 border border-gray-300 rounded-md"
					required
				/>
			</div>

			<!-- Event Date -->
			<div class="mb-4">
				<label for="date" class="block text-gray-700">Date:</label>
				<input
					id="date"
					type="date"
					bind:value={event.date}
					class="w-full p-3 border border-gray-300 rounded-md"
					required
				/>
			</div>
			<!-- Event Time -->
			<div class="mb-4">
				<label for="time" class="block text-gray-700">Time:</label>
				<input
					id="time"
					type="time"
					bind:value={event.time}
					class="w-full p-3 border border-gray-300 rounded-md"
					required
				/>
			</div>
			<!-- Event Short Description -->
			<div class="mb-4">
				<label class="block text-gray-700 text-sm font-bold mb-2" for="short-description"
					>Short Description</label
				>
				<textarea
					id="short-description"
					bind:value={event.short_description}
					class="w-full p-2 border border-gray-300 rounded-md"
					required
				></textarea>
			</div>
			<!-- Event Description -->
			<div class="mb-4">
				<label for="description" class="block text-gray-700">Description:</label>
				<textarea
					id="description"
					bind:value={event.description}
					class="w-full p-3 border border-gray-300 rounded-md"
					required
				></textarea>
			</div>

			<!-- Submit Button -->
			<div class="flex justify-between">
				<button
					type="button"
					class="bg-gray-300 hover:bg-gray-400 text-black py-2 px-4 rounded transition-colors"
					on:click={() => goto(`/events`)}
				>
					Cancel
				</button>
				<button type="submit" class="bg-blue-600 hover:bg-blue-700 text-white py-2 px-4 rounded">
					Save Changes
				</button>
			</div>
		</form>
	</div>
</div>
