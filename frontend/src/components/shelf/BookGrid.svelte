<script lang="ts">
	import BookCard from "./BookCard.svelte";

    let { books, children } = $props();
</script>

<div class="grid grid-flow-col auto-cols-max gap-4">
    {@render children()}
    {#each books as book (book.isbn) }
        <BookCard title={book.title} authors={book.authors} cover={book.cover}>
        </BookCard>
    {/each}
</div>
