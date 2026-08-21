
<script>

	import './layout.css';
	import favicon from '$lib/assets/favicon.svg';
	import Menu from '$lib/components/Menu.svelte';
	import SearchForm from '$lib/components/SearchForm.svelte';

	let { children } = $props();
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>
<SearchForm/>
<Menu/>


<main>
	{@render children()}
</main>

<style>
	nav {
		padding: 1rem;
		background: #f4f4f4;
	}
	main {
		padding: 2rem;
		font-family: sans-serif;
	}
</style>
