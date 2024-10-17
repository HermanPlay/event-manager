<script lang="ts">
	import { onMount } from 'svelte';
	import { browser } from '$app/environment';
	import { addToast } from '$lib/stores/toasts';
	import { getNewUser } from '$lib/utils/utils';
	import type { User, UserInput } from '$lib/schemas/user';
	import UserCard from './components/UserCard.svelte';
	import AddUser from './components/AddUser.svelte';
	import { PUBLIC_API_URL } from '$env/static/public';

	let users: User[] = [];
	let filteredUsers: User[] = [];
	let searchQuery = '';
	let loading = true;
	let errorMessage = '';
	let userToDelete: number;
	let showModal: boolean = false;
	let showAddModal: boolean = false;
	let showEditModal: boolean = false;
	let newUser: UserInput;

	export let data;
	let scrollTop: number = 0;
	let scrollLeft: number = 0;

	function disableScroll() {
		if (browser) {
			scrollTop = window.scrollY || window.document.documentElement.scrollTop;
			scrollLeft = window.scrollX || window.document.documentElement.scrollLeft;
			window.onscroll = function () {
				window.scrollTo(scrollLeft, scrollTop);
			};
		}
	}

	function enableScroll() {
		if (browser) {
			window.onscroll = function () {};
		}
	}

	$: if (showModal || showAddModal) {
		disableScroll();
	} else {
		enableScroll();
	}

	// Fetch users data from an API on component mount
	async function fetchUsers() {
		if (data.error) {
			errorMessage = data.error;
			loading = false;
			return;
		}
		try {
			const response = await fetch(`${PUBLIC_API_URL}/user`, {
				method: 'GET',
				headers: {
					'Content-Type': 'application/json',
					Authorization: 'Bearer ' + data.token
				}
			});
			if (!response.ok) {
				throw new Error('Failed to fetch users.');
			}
			users = (await response.json()).data;
			filteredUsers = users; // Start by displaying all users
		} catch (error: any) {
			addToast({
				type: 'error',
				message: error.message,
				timeout: 5000
			});
		} finally {
			loading = false;
		}
	}

	// Filter users based on the search query
	$: if (searchQuery) {
		filterUsers();
	}

	function filterUsers() {
		if (searchQuery) {
			filteredUsers = users.filter(
				(user) =>
					user.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
					user.email.toLowerCase().includes(searchQuery.toLowerCase()) ||
					user.role.toLowerCase().includes(searchQuery.toLowerCase())
			);
		} else {
			filteredUsers = users;
		}
	}

	// Open the delete confirmation modal and set the user to delete
	function openDeleteModal(userId: number) {
		showModal = true;
		userToDelete = userId;
	}

	// Close the modal without deleting
	function closeModal() {
		showModal = false;
		userToDelete = -1;
	}

	async function confirmDelete() {
		const response = await fetch(`${PUBLIC_API_URL}/user/${userToDelete}`, {
			method: 'DELETE',
			headers: {
				'Content-Type': 'application/json',
				Authorization: 'Bearer ' + data.token
			}
		});
		if (!response.ok) {
			errorMessage = 'Failed to delete user.';
			return;
		}
		users = users.filter((user) => user.id !== userToDelete);
		filterUsers();
		closeModal();
	}

	async function saveNewUser(newUser: UserInput) {
		try {
			const response = await fetch(`${PUBLIC_API_URL}/user`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
					Authorization: 'Bearer ' + data.token
				},
				body: JSON.stringify(newUser)
			});
			if (!response.ok) {
				throw new Error('Failed to save user.');
			}
			const responseData = await response.json();
			// Update the user list
			users = [...users, responseData.data];
			filteredUsers = users;
			showAddModal = false;
		} catch (error: any) {
			errorMessage = error.message;
		}
	}

	async function saveEditedUser(userToEdit: User) {
		try {
			const response = await fetch(`${PUBLIC_API_URL}/user/${userToEdit.id}`, {
				method: 'PATCH',
				headers: {
					'Content-Type': 'application/json',
					Authorization: 'Bearer ' + data.token
				},
				body: JSON.stringify(userToEdit)
			});
			if (!response.ok) {
				throw new Error('Failed to update user.');
			}

			// Update the user list
			users.filter((user) => (user.id === userToEdit.id ? { ...userToEdit } : user));
			filteredUsers = users;
			showEditModal = false;
		} catch (error: any) {
			errorMessage = error.message;
		}
	}

	async function closeAddModal() {
		showAddModal = false;
	}

	onMount(fetchUsers);
</script>

<!-- Page Layout -->
<div class="min-h-screen p-8 bg-gray-100">
	<div class="container mx-auto bg-white shadow-lg rounded-md p-6">
		<h1 class="text-3xl font-bold mb-6">Users</h1>

		<!-- Tabs -->
		<div class="mb-4 flex justify-between w-full">
			<button
				class={`px-4 py-2 rounded-md bg-blue-500 text-white`}
				on:click={() => {
					showAddModal = true;
					newUser = getNewUser();
				}}
			>
				+ Add user
			</button>
		</div>

		<!-- Search Input -->
		<div class="mb-4">
			<input
				type="text"
				placeholder="Search user..."
				bind:value={searchQuery}
				class="w-full p-3 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
			/>
		</div>

		<!-- Loading Spinner -->
		{#if loading}
			<div class="flex justify-center items-center py-8">
				<span class="animate-spin text-blue-500 text-3xl">⏳</span>
			</div>
		{:else if errorMessage}
			<p class="text-red-500 text-center">{errorMessage}</p>
		{:else}
			<!-- User List -->
			{#if filteredUsers.length > 0}
				<ul>
					{#each filteredUsers as user}
						<li class="mb-6 p-4 bg-gray-50 shadow-sm rounded-lg relative">
							<UserCard
								{user}
								onEdit={() => {
									showEditModal = true;
									newUser = user;
								}}
								onDelete={openDeleteModal}
							/>
						</li>
					{/each}
				</ul>
			{:else}
				<p class="text-center text-gray-500">No users found.</p>
			{/if}
		{/if}
	</div>

	{#if showAddModal}
		<AddUser saveUser={saveNewUser} user={newUser} closeModal={closeAddModal} />
	{/if}
	{#if showEditModal}
		<AddUser
			saveUser={saveEditedUser}
			user={newUser}
			closeModal={() => {
				showEditModal = false;
			}}
		/>
	{/if}

	<!-- Delete Confirmation Modal -->
	{#if showModal}
		<div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center">
			<div class="bg-white p-6 rounded-lg shadow-lg max-w-sm w-full">
				<h2 class="text-xl font-semibold mb-4">Are you sure?</h2>
				<p class="text-gray-600 mb-6">
					Do you really want to delete this user? This action cannot be undone.
				</p>
				<div class="flex justify-end space-x-4">
					<button
						class="bg-gray-300 hover:bg-gray-400 text-gray-800 py-2 px-4 rounded"
						on:click={closeModal}
					>
						Cancel
					</button>
					<button
						class="bg-red-600 hover:bg-red-700 text-white py-2 px-4 rounded"
						on:click={confirmDelete}
					>
						Delete
					</button>
				</div>
			</div>
		</div>
	{/if}
</div>
