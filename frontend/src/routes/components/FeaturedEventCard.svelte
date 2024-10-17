<script lang="ts">
	import type { Event } from '$lib/schemas/event';

	export let event: Event;
</script>

<div class="bg-white shadow-lg rounded-lg p-4 hover:shadow-xl">
	<h3 class="font-semibold text-xl mb-2 hover:underline">
		<a href={`/events/${event.id}`}>
			{event.title}
		</a>
	</h3>

	<div class="flex gap-2">
		<p class="text-gray-600">{event.date}</p>
		<p class="text-gray-600">{event.location}</p>
	</div>
	<p class="text-gray-600">{event.short_description}</p>
</div>
