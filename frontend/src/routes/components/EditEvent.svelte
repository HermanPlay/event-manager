<script lang="ts">
	export let saveEvent;
	export let event;
	export let closeEditModal;
</script>

<div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center">
	<div class="bg-white p-6 rounded-lg shadow-lg max-w-lg w-full">
		<h2 class="text-xl font-semibold mb-4">Edit Event</h2>

		<!-- Event Edit Form -->
		<form on:submit|preventDefault={() => saveEvent(event)}>
			<div class="mb-4">
				<label class="block text-gray-700 text-sm font-bold mb-2" for="title">Title</label>
				<input
					type="text"
					id="title"
					bind:value={event.title}
					class="w-full p-2 border border-gray-300 rounded-md"
					required
				/>
			</div>
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
			<div class="mb-4">
				<label class="block text-gray-700 text-sm font-bold mb-2" for="description"
					>Description</label
				>
				<textarea
					id="description"
					bind:value={event.description}
					class="w-full p-2 border border-gray-300 rounded-md"
					required
				></textarea>
			</div>
			<div class="mb-4">
				<label class="block text-gray-700 text-sm font-bold mb-2" for="date">Date</label>
				<input
					type="date"
					id="date"
					bind:value={event.date}
					class="w-full p-2 border border-gray-300 rounded-md"
					required
				/>
			</div>
			<div class="mb-4">
				<label class="block text-gray-700 text-sm font-bold mb-2" for="location">Location</label>
				<input
					type="text"
					id="location"
					bind:value={event.location}
					class="w-full p-2 border border-gray-300 rounded-md"
					required
				/>
			</div>
			<div class="flex justify-end space-x-4">
				<button
					type="button"
					class="bg-gray-300 hover:bg-gray-400 text-gray-800 py-2 px-4 rounded"
					on:click={closeEditModal}
				>
					Cancel
				</button>
				<button type="submit" class="bg-blue-600 hover:bg-blue-700 text-white py-2 px-4 rounded">
					Save
				</button>
			</div>
		</form>
	</div>
</div>
