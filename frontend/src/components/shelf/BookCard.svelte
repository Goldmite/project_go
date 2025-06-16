<script lang="ts">
	let { title, authors, cover } = $props();
</script>

<div class="w-42 sm:w-56">
	<div class="h-60 sm:h-84 mb-2">
		<img class="h-full rounded-md border object-cover" src={cover} alt={title} />
	</div>
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
