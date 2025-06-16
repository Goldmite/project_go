<script lang="ts">
	import BookCard from './BookCard.svelte';

	let { books, children } = $props();
</script>

<div class="grid justify-start gap-4
	grid-cols-[repeat(auto-fit,minmax(10.5rem,auto))]
	sm:grid-cols-[repeat(auto-fit,minmax(14rem,auto))]"
>
	{@render children()}
	{#each books as book (book.isbn)}
		<BookCard title={book.title} authors={book.authors} cover={book.cover}></BookCard>
	{/each}
</div>
