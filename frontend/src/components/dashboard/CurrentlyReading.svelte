<script lang="ts">
	import BookCard from '../shelf/BookCard.svelte';
	import Grid from './Grid.svelte';

	let { children } = $props();
</script>

<div class="ml-4 sm:ml-10">
	{@render children()}
	<Grid>
		<BookCard ownedBy={[]} isbn="11111111111" title="Test" authors="Placeholder" cover="" />
	</Grid>
</div>

<style>
</style>
