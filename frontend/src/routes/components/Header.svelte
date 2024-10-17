<!-- Inspired by https://codepen.io/syahrizaldev/pen/WNmMmbK -->
<script lang="ts">
	import { isAuthenticated } from '$lib/stores/user';

	let is_active: boolean = false;
	let innerWidth: number = 0;

	$: if (innerWidth > 768) {
		is_active = true;
	}

	const nav_items = [
		{
			name: 'Home',
			href: '/'
		},
		{
			name: 'Events',
			href: '/events'
		},
		{
			name: 'About Us',
			href: '/about'
		},
		{
			name: 'Contact',
			href: '/contact'
		}
	];

	const toggleMenu = () => {
		if (innerWidth <= 768) {
			is_active = !is_active;
		}
	};
</script>

<svelte:window bind:innerWidth />
<header>
	<nav
		class="w-full h-16 ms-auto me-auto md:flex md:justify-between max-w-5xl px-6 text-gray-200 bg-gray-800"
	>
		<div
			class="flex items-center justify-between max-w-full h-full md:opacity-1 md:pointer-events-auto md:bg-none md:transition-none"
		>
			<a class="text-2xl font-bold leading-tight text-white text-nowrap" href="/">Event Manager</a>
			<div
				class="relative block cursor-pointer select-none w-6 h-4 border-none outline-none visible md:hidden md:invisible"
			>
				<button on:click={toggleMenu} class="w-7 h-5">
					<span
						class="burger-line"
						style={is_active ? 'top: 0.5rem; rotate: 135deg' : 'top: 0; rotate:0deg'}
					></span>
					<span class="burger-line" style={is_active ? 'opacity: 0' : 'top: 0.5rem'}></span>
					<span class="burger-line" style={is_active ? 'top: 0.5rem; rotate: -135deg' : 'top: 1rem'}
					></span>
				</button>
			</div>
		</div>
		<div
			class="nav-block absolute w-full h-[calc(100vh - 4rem)] opacity-{is_active
				? '1'
				: '0'} overflow-hidden {is_active
				? 'flex justify-center'
				: 'hidden'} transition-opacity duration-400 ease-in-out md:transition-none md:opacity-1 bg-gray-800"
		>
			<ul class="py-4 md:flex md:flex-row md:justify-end md:gap-1 md:h-full md:p-0 md:bg-gray-800">
				{#each nav_items as item}
					<li class="list-item md:flex md:cursor-pointer md:px-2 text-pretty break-words">
						<a
							on:click={toggleMenu}
							class="menu-link text-base font-medium flex items-center justify-between md:p-0 transition-colors duration-300 ease-in-out"
							href={item.href}>{item.name}</a
						>
					</li>
				{/each}
				{#if $isAuthenticated}
					<li class="list-item md:flex md:cursor-pointer md:px-2 text-pretty break-words">
						<a
							on:click={toggleMenu}
							class="text-center menu-link text-base font-normal underline flex items-center justify-between md:p-0 transition-colors duration-300 ease-in-out"
							href="/signout">Sign out</a
						>
					</li>
				{:else}
					<li class="list-item md:flex md:cursor-pointer md:px-2 text-pretty break-words">
						<a
							on:click={toggleMenu}
							class="menu-link text-base font-normal underline flex items-center justify-between md:p-0 transition-colors duration-300 ease-in-out"
							href="/auth">Sign in</a
						>
					</li>
				{/if}
			</ul>
		</div>
	</nav>
</header>

<style>
	.burger-line {
		position: absolute;
		display: block;
		right: 0;
		width: 100%;
		height: 2.15px;
		opacity: 1;
		border-radius: 0.15rem;
		background-color: white;
		transition: all 0.3s ease;
	}

	@media screen and (min-width: 768px) {
		.nav-block {
			position: initial;
			height: initial;
			pointer-events: visible;
		}
	}
	.menu-link {
		padding-block: 0.5rem;
		padding-inline: 1rem;
	}
</style>
