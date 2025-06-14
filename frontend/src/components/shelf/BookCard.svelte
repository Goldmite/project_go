<script lang="ts">
	let { title, authors, cover } = $props();
</script>

<div class="card">
	<div class="imageContainer">
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
	.card {
		width: 14rem;
	}
	.imageContainer {
		height: 20rem;
		margin-bottom: 0.5em;
	}
	p {
		font-family: sans-serif;
		text-wrap: wrap;
	}
</style>
