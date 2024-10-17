<script lang="ts">
	export let toggle: boolean = false;
</script>

<div>
	<input
		class="opacity-0 block absolute h-8 w-8 z-10 cursor-pointer"
		type="checkbox"
		bind:checked={toggle}
	/>
	<span></span>
	<span></span>
	<span></span>
</div>

<style>
	/*
 * Just a quick hamburger
 */
	span {
		display: block;
		width: 33px;
		height: 4px;
		margin-bottom: 5px;
		position: relative;

		background: #cdcdcd;
		border-radius: 3px;

		z-index: 1;

		transform-origin: 4px 0px;

		transition:
			transform 0.5s cubic-bezier(0.77, 0.2, 0.05, 1),
			background 0.5s cubic-bezier(0.77, 0.2, 0.05, 1),
			opacity 0.55s ease;
	}

	span:first-child {
		transform-origin: 0% 0%;
	}

	span:nth-last-child(1) {
		transform-origin: 0% 100%;
		margin-bottom: 0;
	}

	input:checked ~ span:nth-last-child(3) {
		transform: rotate(45deg) translate(-2px, 0px);
	}
	input:checked ~ span:nth-last-child(2) {
		opacity: 0;
		transform: rotate(0deg) scale(0.2, 0.2);
	}
	input:checked ~ span:nth-last-child(1) {
		transform: rotate(-45deg) translate(0px, 0px);
	}
</style>
