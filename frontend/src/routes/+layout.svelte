<script lang="ts">
	import '../app.css';
	import Footer from './components/Footer.svelte';
	import Toasts from './components/Toasts/Toasts.svelte';
	import Header from './components/Header.svelte';
</script>

<Toasts />
<Header />
<slot />
<Footer />
