<script lang="ts">
	import { onMount } from 'svelte';
	import { browser } from '$app/environment';
	import type { Event } from '$lib/schemas/event';
	import EventCard from './components/EventCard.svelte';
	import { addToast } from '$lib/stores/toasts';
	import { goto } from '$app/navigation';
	import { PUBLIC_API_URL } from '$env/static/public';

	let events: Event[] = []; // Array to store all events
	let filteredEvents: Event[] = []; // Array to store the filtered events based on search and tabs
	let myEvents: Event[] = []; // Array to store events created or booked by the current user
	let searchQuery = '';
	let loading = true;
	let errorMessage = '';
	let showModal: boolean = false;
	let eventToDelete: number;

	export let data;
	let scrollTop: number = 0;
	let scrollLeft: number = 0;

	// Tabs state
	let activeTab: 'all' | 'my' = 'all'; // Default tab is "All Events"

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

	$: if (showModal) {
		disableScroll();
	} else {
		enableScroll();
	}

	// Display controls logic for edit/delete actions
	function displayControls(event: Event) {
		return data.currentUserID === event.created_by || data.userRole === 'admin';
	}

	function isAdmin() {
		return data.userRole === 'admin';
	}

	// Fetch events data from an API on component mount
	async function fetchEvents() {
		if (data.error) {
			errorMessage = data.error;
			loading = false;
			return;
		}
		// All events
		try {
			const response = await fetch(`${PUBLIC_API_URL}/event`, {
				method: 'GET',
				headers: {
					'Content-Type': 'application/json',
					Authorization: 'Bearer ' + data.token
				}
			});
			if (!response.ok) {
				throw new Error('Failed to fetch events.');
			}
			events = (await response.json()).data;
			filteredEvents = events; // Start by displaying all events
		} catch (error: any) {
			addToast({
				type: 'error',
				message: error.message,
				timeout: 5000
			});
		}
		// MyEvents
		try {
			const respose = await fetch(`${PUBLIC_API_URL}/event/my/${data.currentUserID}`, {
				method: 'GET',
				headers: {
					'Content-Type': 'application/json',
					Authorization: 'Bearer ' + data.token
				}
			});

			if (!respose.ok) {
				throw new Error('Failed to fetch my events.');
			}

			myEvents = (await respose.json()).data;
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

	// Filter events based on the active tab and search query
	$: if (searchQuery || activeTab) {
		filterEvents();
	}

	function filterEvents() {
		let eventList = activeTab === 'my' ? myEvents : events; // Choose event list based on tab
		if (searchQuery) {
			filteredEvents = eventList.filter(
				(event) =>
					event.title.toLowerCase().includes(searchQuery.toLowerCase()) ||
					event.description.toLowerCase().includes(searchQuery.toLowerCase()) ||
					event.location.toLowerCase().includes(searchQuery.toLowerCase())
			);
		} else {
			filteredEvents = eventList;
		}
	}

	// Open the delete confirmation modal and set the event to delete
	function openDeleteModal(eventId: number) {
		showModal = true;
		eventToDelete = eventId;
	}

	// Close the modal without deleting
	function closeModal() {
		showModal = false;
		eventToDelete = -1;
	}

	// Delete the event after confirmation
	async function confirmDelete() {
		const response = await fetch(`${PUBLIC_API_URL}/event/${eventToDelete}`, {
			method: 'DELETE',
			headers: {
				'Content-Type': 'application/json',
				Authorization: 'Bearer ' + data.token
			}
		});
		if (!response.ok) {
			errorMessage = 'Failed to delete event.';
			return;
		}
		events = events.filter((event) => event.id !== eventToDelete);
		myEvents = myEvents.filter((event) => event.id !== eventToDelete);
		filterEvents();
		closeModal();
	}

	async function changeFeature(event: Event) {
		const response = await fetch(`${PUBLIC_API_URL}/event/${event.id}`, {
			method: 'PATCH',
			headers: {
				'Content-Type': 'application/json',
				Authorization: 'Bearer ' + data.token
			},

			body: JSON.stringify({ is_featured: event.is_featured })
		});
		if (!response.ok) {
			addToast({
				type: 'error',
				message: 'Failed to change event feature.',
				timeout: 5000
			});
			console.error(await response.json());
		}
		addToast({
			type: 'success',
			message: 'Event feature updated successfully.',
			timeout: 2000
		});
	}

	async function bookEvent(eventId: number) {
		try {
			const response = await fetch(`${PUBLIC_API_URL}/event/book/${eventId}`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
					Authorization: 'Bearer ' + data.token
				}
			});
			if (!response.ok) {
				const data = await response.json();
				if (response.status === 409) {
					addToast({
						type: 'info',
						message: "You've already booked this event.",
						timeout: 5000
					});
				} else {
					console.log(data);
					throw new Error('Failed to book event.');
				}
			}
		} catch (error: any) {
			addToast({
				type: 'error',
				message: error.message,
				timeout: 5000
			});
		}
		await fetchEvents();
	}

	onMount(fetchEvents);
</script>

<!-- Page Layout -->
<div class="min-h-screen p-8 bg-gray-100">
	<div class="container mx-auto bg-white shadow-lg rounded-md p-6">
		<h1 class="text-3xl font-bold mb-6">Events</h1>

		<!-- Tabs -->
		<div class="mb-4 flex justify-between w-full">
			<div>
				<button
					class={`px-4 py-2 rounded-md ${activeTab === 'all' ? 'bg-blue-500 text-white' : 'bg-gray-200 text-gray-800'}`}
					on:click={() => {
						activeTab = 'all';
						filterEvents();
					}}
				>
					All Events
				</button>
				<button
					class={`px-4 py-2 rounded-md ${activeTab === 'my' ? 'bg-blue-500 text-white' : 'bg-gray-200 text-gray-800'}`}
					on:click={() => {
						activeTab = 'my';
						filterEvents();
					}}
				>
					My Events
				</button>
			</div>

			<button
				class="px-4 py-2 rounded-md bg-blue-500 text-white hover:bg-blue-600"
				on:click={() => goto('/events/new')}
			>
				+ Add event
			</button>
		</div>

		<!-- Search Input -->
		<div class="mb-4">
			<input
				type="text"
				placeholder="Search events..."
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
			<!-- Event List -->
			{#if filteredEvents.length > 0}
				<ul>
					{#each filteredEvents as event}
						<li class="mb-6 p-4 shadow-sm rounded-lg relative">
							<EventCard
								{event}
								onEdit={() => goto(`/events/edit/${event.id}`)}
								onDelete={openDeleteModal}
								onChangeFeature={changeFeature}
								doDisplayControls={displayControls}
								doDisplayFeature={isAdmin}
								{bookEvent}
							/>
						</li>
					{/each}
				</ul>
			{:else}
				<p class="text-center text-gray-500">No events found.</p>
			{/if}
		{/if}
	</div>

	<!-- Delete Confirmation Modal -->
	{#if showModal}
		<div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center">
			<div class="bg-white p-6 rounded-lg shadow-lg max-w-sm w-full">
				<h2 class="text-xl font-semibold mb-4">Are you sure?</h2>
				<p class="text-gray-600 mb-6">
					Do you really want to delete this event? This action cannot be undone.
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
