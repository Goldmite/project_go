<script lang="ts">
	import BookCard from './BookCard.svelte';

	let { books, children = () => {} } = $props();
</script>

<div
	class="grid grid-cols-[repeat(auto-fit,minmax(10.5rem,auto))] justify-start
	gap-4
	sm:grid-cols-[repeat(auto-fit,minmax(14rem,auto))]"
>
	{@render children()}
	{#each books as book (book.isbn)}
		<BookCard isbn={book.isbn} title={book.title} authors={book.authors} cover={book.cover} />
	{/each}
</div>
