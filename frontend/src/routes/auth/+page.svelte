<script lang="ts">
	import { goto } from '$app/navigation';
	import { addToast } from '$lib/stores/toasts';
	import { isAuthenticated } from '$lib/stores/user';
	import { PUBLIC_API_URL } from '$env/static/public';

	let isResettingPassword: boolean = false; // Track whether to show the reset password form
	let isRegistering: boolean = true; // Track whether to show the registration or login form
	let name: string = '';
	let email: string = '';
	let password: string = '';
	let confirmPassword: string = '';
	let formError: string = '';
	let isSubmitting: boolean = false;

	async function handleRegister() {
		formError = '';

		// Validate registration form
		if (!name || !email || !password || !confirmPassword) {
			formError = 'All fields are required';
			return;
		}

		if (password !== confirmPassword) {
			formError = 'Passwords do not match';
			return;
		}

		isSubmitting = true;

		// Send registration request to API
		try {
			const response = await fetch(`${PUBLIC_API_URL}/auth/register`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ name, email, password })
			});

			if (response.ok) {
				addToast({
					type: 'success',
					message: 'Registration successful! You can now log in.',
					timeout: 5000
				});
				isRegistering = false; // Switch to login form after successful registration
			} else {
				const data = await response.json();
				formError = data.message || 'Registration failed. Try again.';
			}
		} catch (error: any) {
			formError = `${error.message}. Please try again later.`;
		} finally {
			isSubmitting = false;
		}
	}

	async function handleLogin() {
		formError = '';

		// Validate login form
		if (!email || !password) {
			formError = 'Email and password are required';
			return;
		}

		isSubmitting = true;

		// Send login request to API
		try {
			const response = await fetch(`${PUBLIC_API_URL}/auth/login`, {
				method: 'POST',
				credentials: 'include',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ email, password })
			});

			if (response.ok) {
				addToast({
					type: 'success',
					message: 'Login successful!',
					timeout: 5000
				});

				isAuthenticated.set(true);

				await goto('/', { invalidateAll: true });
			} else {
				const data = await response.json();
				formError = data.data.message || 'Login failed. Try again.';
			}
		} catch (error: any) {
			formError = `${error.message}. Please try again later.`;
		} finally {
			isSubmitting = false;
		}
	}

	// Password Reset Handler
	async function handlePasswordReset() {
		formError = '';

		if (!email) {
			formError = 'Email is required';
			return;
		}

		isSubmitting = true;

		try {
			const response = await fetch(`${PUBLIC_API_URL}/auth/reset`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ email })
			});

			if (response.ok) {
				const data = await response.json();
				alert('Your new password is: ' + data.data.new_password);
				isRegistering = false; // Switch back to login after reset request
			} else {
				const data = await response.json();
				formError = data.message || 'Password reset failed. Try again.';
			}
		} catch (error: any) {
			formError = `${error.message}. Please try again later.`;
		} finally {
			isSubmitting = false;
		}
	}

	function toggleForm(isReg: boolean) {
		isRegistering = isReg; // Toggle between registration and login forms
	}
	function showResetPasswordForm() {
		isResettingPassword = true;
		isRegistering = false;
	}

	function showLoginForm() {
		isResettingPassword = false;
		isRegistering = false;
	}
</script>

<div class="min-h-screen flex items-center justify-center bg-gray-100">
	<div class="bg-white p-8 shadow-lg rounded-md max-w-md w-full">
		<h1 class="text-2xl font-bold mb-6">
			{#if isResettingPassword}
				Reset Password
			{:else if isRegistering}
				Register
			{:else}
				Login
			{/if}
		</h1>
		{#if formError}
			<p class="text-red-500 mb-4">{formError}</p>
		{/if}

		{#if !isResettingPassword}
			<!-- Button to choose between Register and Login -->
			<div class="flex justify-between mb-6">
				<button
					class="w-full bg-blue-600 text-white p-3 rounded mr-2 transition duration-200 transform hover:bg-blue-700 hover:scale-105"
					on:click={() => toggleForm(true)}
					class:opacity-50={!isRegistering}
					disabled={isSubmitting}
				>
					Register
				</button>
				<button
					class="w-full bg-blue-600 text-white p-3 rounded ml-2 transition duration-200 transform hover:bg-blue-700 hover:scale-105"
					on:click={() => toggleForm(false)}
					class:opacity-50={isRegistering}
					disabled={isSubmitting}
				>
					Login
				</button>
			</div>
		{/if}

		{#if isResettingPassword}
			<!-- Password Reset Form -->
			<form on:submit|preventDefault={handlePasswordReset}>
				<!-- Email Field -->
				<div class="mb-4">
					<label for="email" class="block font-medium">Email</label>
					<input
						type="email"
						id="email"
						bind:value={email}
						required
						class="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500 transition duration-200"
						placeholder="Enter your email"
					/>
				</div>

				<!-- Submit Button -->
				<button
					type="submit"
					class="w-full bg-blue-600 text-white p-3 rounded transition duration-200 hover:bg-blue-700"
					disabled={isSubmitting}
				>
					{isSubmitting ? 'Resetting...' : 'Reset Password'}
				</button>
			</form>

			<p class="mt-4 text-sm">
				Remembered your password?
				<a class="text-blue-600 hover:underline cursor-pointer" on:click={showLoginForm}>
					Log in
				</a>
			</p>
		{:else if isRegistering}
			<div class="transition-opacity duration-300" class:opacity={isRegistering ? 1 : 0}>
				<form on:submit|preventDefault={handleRegister}>
					<!-- Name Field (only for registration) -->
					<div class="mb-4">
						<label for="name" class="block font-medium">Name</label>
						<input
							type="text"
							id="name"
							bind:value={name}
							required
							class="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500 transition duration-200"
							placeholder="Enter your name"
						/>
					</div>

					<!-- Email Field -->
					<div class="mb-4">
						<label for="email" class="block font-medium">Email</label>
						<input
							type="email"
							id="email"
							bind:value={email}
							required
							class="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500 transition duration-200"
							placeholder="Enter your email"
						/>
					</div>

					<!-- Password Field -->
					<div class="mb-4">
						<label for="password" class="block font-medium">Password</label>
						<input
							type="password"
							id="password"
							bind:value={password}
							required
							class="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500 transition duration-200"
							placeholder="Enter your password"
						/>
					</div>

					<!-- Confirm Password Field -->
					<div class="mb-4">
						<label for="confirmPassword" class="block font-medium">Confirm Password</label>
						<input
							type="password"
							id="confirmPassword"
							bind:value={confirmPassword}
							required
							class="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500 transition duration-200"
							placeholder="Confirm your password"
						/>
					</div>

					<!-- Submit Button -->
					<button
						type="submit"
						class="w-full bg-blue-600 text-white p-3 rounded transition duration-200 hover:bg-blue-700"
						disabled={isSubmitting}
					>
						{isSubmitting ? 'Registering...' : 'Register'}
					</button>
				</form>
			</div>
		{:else}
			<div class="transition-opacity duration-300" class:opacity={isRegistering ? 0 : 1}>
				<form on:submit|preventDefault={handleLogin}>
					<!-- Email Field -->
					<div class="mb-4">
						<label for="email" class="block font-medium">Email</label>
						<input
							type="email"
							id="email"
							bind:value={email}
							required
							class="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500 transition duration-200"
							placeholder="Enter your email"
						/>
					</div>

					<!-- Password Field -->
					<div class="mb-4">
						<label for="password" class="block font-medium">Password</label>
						<input
							type="password"
							id="password"
							bind:value={password}
							required
							class="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500 transition duration-200"
							placeholder="Enter your password"
						/>
					</div>

					<!-- Submit Button -->
					<button
						type="submit"
						class="w-full bg-blue-600 text-white p-3 rounded transition duration-200 hover:bg-blue-700"
						disabled={isSubmitting}
					>
						{isSubmitting ? 'Logging in...' : 'Login'}
					</button>
				</form>
				<p class="mt-4 text-sm">
					Forgot your password?
					<a class="text-blue-600 hover:underline cursor-pointer" on:click={showResetPasswordForm}>
						Reset it here
					</a>
				</p>
			</div>
		{/if}
	</div>
</div>
