<script lang="ts">
	import { addToast } from '$lib/stores/toasts';

	let name: string = '';
	let email: string = '';
	let message: string = '';
	let isSubmitting: boolean = false;
	let formResponse: string = '';

	// Form Submission Handler
	async function handleSubmit() {
		isSubmitting = true;
		formResponse = '';

		// Construct form data object
		const formData = {
			name,
			email,
			message
		};

		try {
			addToast({
				message: 'Form submitted successfully!',
				type: 'success'
			});
		} catch (error: any) {
			formResponse = `${error.message}. Please try again later.`;
		} finally {
			isSubmitting = false;
			resetForm();
		}
	}

	// Reset form fields after submission
	function resetForm() {
		name = '';
		email = '';
		message = '';
	}
</script>

<!-- Contact Page -->
<div class="min-h-screen bg-gray-100 p-10">
	<div class="container mx-auto max-w-2xl bg-white p-8 shadow-md rounded">
		<h1 class="text-3xl font-bold mb-6">Contact Us</h1>
		<p class="text-gray-600 mb-6">Feel free to reach out if you have any questions or feedback.</p>

		<!-- Contact Form -->
		<form on:submit|preventDefault={handleSubmit}>
			<!-- Name Field -->
			<div class="mb-4">
				<label for="name" class="block text-lg font-medium">Your Name</label>
				<input
					type="text"
					id="name"
					bind:value={name}
					required
					class="w-full p-2 border border-gray-300 rounded text-black"
					placeholder="Enter your name"
				/>
			</div>

			<!-- Email Field -->
			<div class="mb-4">
				<label for="email" class="block text-lg font-medium">Your Email</label>
				<input
					type="email"
					id="email"
					bind:value={email}
					required
					class="w-full p-2 border border-gray-300 rounded text-black"
					placeholder="Enter your email"
				/>
			</div>

			<!-- Message Field -->
			<div class="mb-4">
				<label for="message" class="block text-lg font-medium">Your Message</label>
				<textarea
					id="message"
					bind:value={message}
					required
					class="w-full p-2 border border-gray-300 rounded text-black"
					placeholder="Enter your message"
					rows="5"
				></textarea>
			</div>

			<!-- Submit Button -->
			<button
				type="submit"
				class="w-full bg-blue-600 text-white p-3 rounded mt-4"
				disabled={isSubmitting}
			>
				{isSubmitting ? 'Submitting...' : 'Submit'}
			</button>
		</form>

		<!-- Response Message -->
		{#if formResponse}
			<p class="mt-6 text-center text-lg font-semibold text-green-500">
				{formResponse}
			</p>
		{/if}
	</div>
</div>

<style>
	/* Custom styles if necessary */
</style>
