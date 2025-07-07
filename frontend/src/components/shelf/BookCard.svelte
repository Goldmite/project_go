<script lang="ts">
	import { bookOwners } from '$lib/stores/localMembers';
	import CoverWrapper from './CoverWrapper.svelte';

	let { ownedBy, isbn, title, authors, cover } = $props();

</script>

<div class="w-42 sm:w-56">
	<a data-sveltekit-preload-data="tap" onclick={() => bookOwners.set(ownedBy)} href="/book/{isbn}">
		<CoverWrapper {cover} {title} />
	</a>
	<p class="font-semibold">{title}</p>
	<p class="line-clamp-2 text-sm font-light italic">
		{#each authors as author, i}
			{author}{i + 1 < authors.length ? ', ' : ''}
		{/each}
	</p>
</div>

<style>
	p {
		font-family: sans-serif;
		text-wrap: wrap;
	}
</style>
