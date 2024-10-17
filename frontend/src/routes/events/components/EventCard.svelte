<script lang="ts">
	import type { Event } from '$lib/schemas/event';

	export let event: Event;
	export let onEdit: (event: Event) => void;
	export let onDelete: (id: number) => void;
	export let onChangeFeature: (event: Event) => void;
	export let doDisplayControls: (event: Event) => boolean;
	export let doDisplayFeature: () => boolean;
	export let bookEvent: (eventId: number) => void;
	$: featureText = event.is_featured ? 'UnFeature event' : 'Feature event';
</script>

<div id="event-card" class="p-4 bg-white shadow-md rounded-md relative">
	<h2 class="text-xl font-semibold mb-2 hover:underline">
		<a href={`/events/${event.id}`}>{event.title}</a>
	</h2>
	<div class="flex gap-2">
		<p class="text-gray-600">{event.date} {event.time},</p>
		<p class="text-gray-600">{event.location}</p>
	</div>
	<p class="text-gray-700 mt-2">
		{event.short_description}
	</p>

	<!-- Display controls based on displayControls() -->
	<div
		class="buttons-container mt-4 space-x-2 md:absolute md:top-4 md:right-4 md:space-x-2 md:space-y-0 flex justify-center"
	>
		{#if doDisplayControls(event)}
			{#if doDisplayFeature()}
				<button
					class="text-white bg-blue-600 p-2 rounded hover:bg-blue-700 w-full md:w-auto"
					on:click={() => {
						event.is_featured = !event.is_featured;
						onChangeFeature(event);
					}}
				>
					{featureText}
				</button>
			{/if}
			<button
				class="text-white bg-blue-600 p-2 rounded hover:bg-blue-700 w-full md:w-auto"
				on:click={() => onEdit(event)}
			>
				Edit
			</button>
			<button
				class="text-white bg-red-600 p-2 rounded hover:bg-red-700 w-full md:w-auto"
				on:click={() => onDelete(event.id)}
			>
				Delete
			</button>
		{:else}
			<button
				class="text-white bg-blue-600 p-2 rounded hover:bg-blue-700 w-full md:w-auto"
				on:click={() => bookEvent(event.id)}
			>
				Book
			</button>
		{/if}
	</div>
</div>

<style>
	#event-card {
		position: relative;
	}

	@media (min-width: 768px) {
		/* On larger screens, position the buttons on the side */
		.buttons-container {
			flex-direction: row;
		}
	}
</style>
