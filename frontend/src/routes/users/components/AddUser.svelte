<script lang="ts">
	export let saveUser;
	export let user;
	export let closeModal;
</script>

<div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center">
	<div class="bg-white p-6 rounded-lg shadow-lg max-w-lg w-full">
		<h2 class="text-xl font-semibold mb-4">Edit Event</h2>

		<!-- Event Edit Form -->
		<form on:submit|preventDefault={() => saveUser(user)}>
			<div class="mb-4">
				<label class="block text-gray-700 text-sm font-bold mb-2" for="title">Name</label>
				<input
					type="text"
					id="name"
					bind:value={user.name}
					class="w-full p-2 border border-gray-300 rounded-md"
					required
				/>
			</div>
			<div class="mb-4">
				<label class="block text-gray-700 text-sm font-bold mb-2" for="description">Email</label>
				<input
					type="email"
					id="email"
					bind:value={user.email}
					class="w-full p-2 border border-gray-300 rounded-md"
					required
				/>
			</div>
			<div class="mb-4">
				<label class="block text-gray-700 text-sm font-bold mb-2" for="role">Role</label>
				<select
					bind:value={user.role}
					class="w-full p-2 border border-gray-300 rounded-md"
					required
				>
					<option value="user">User</option>
					<option value="manager">Manager</option>
					<option value="admin">Admin</option>
				</select>
			</div>
			<div class="flex justify-end space-x-4">
				<button
					type="button"
					class="bg-gray-300 hover:bg-gray-400 text-gray-800 py-2 px-4 rounded"
					on:click={closeModal}
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
