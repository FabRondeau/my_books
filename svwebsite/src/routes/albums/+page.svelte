<script>

	import { error } from '@sveltejs/kit';
import { onMount } from 'svelte';

	let url = 'http://localhost:3000/albums';
	let albums = [];

	// onMount(async () => {
		// albums = await getAlbums();
	// });

	const getAlbums = async () => {
		const res = await fetch(url);
		const data = await res.json();
		console.log(data);
		// const filteredData = data.slice(0, 3);
		// console.log(filteredData);
		return data;
	};
	let albumsPromise = getAlbums();
	console.log(albumsPromise)

</script>

<h1>Nos albums</h1>
{#await albumsPromise}
<p>Chargement des albums...</p>
{:then albums}
<table>
	<thead>
		<tr>
			<th>Titre</th>
			<th>Artiste</th>
			<th>Prix</th></tr>
	</thead>
	<tbody>
	{#each albums as album}
	<tr>
		<td>{album.title}</td>
		<td>{album.artist}</td>
		<td>{album.price}</td>
	</tr>
	{/each}
	</tbody>
</table>
	{:catch error}
	<p>Erreur: {error.message}</p>
{/await}
