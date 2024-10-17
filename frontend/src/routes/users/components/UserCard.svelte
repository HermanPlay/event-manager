<script lang="ts">
	import type { User } from '$lib/schemas/user';

	export let user: User;
	export let onEdit: (event: User) => void;
	export let onDelete: (id: number) => void;
</script>

<div id="event-card" class="p-4 bg-white shadow-md rounded-md relative">
	<h2 class="text-xl font-semibold mb-2">{user.name}</h2>
	<div class="flex gap-2">
		<p class="text-gray-600">{user.email}</p>
		<p class="text-gray-600">{user.role}</p>
	</div>

	<!-- Display controls based on displayControls() -->
	<div
		class="buttons-container mt-4 space-x-2 md:absolute md:top-4 md:right-4 md:space-x-2 md:space-y-0 flex justify-center"
	>
		<button
			class="text-white bg-blue-600 p-2 rounded hover:bg-blue-700 w-full md:w-auto"
			on:click={() => onEdit(user)}
		>
			Edit
		</button>
		<button
			class="text-white bg-red-600 p-2 rounded hover:bg-red-700 w-full md:w-auto"
			on:click={() => {
				onDelete(user.id);
			}}
		>
			Delete
		</button>
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
