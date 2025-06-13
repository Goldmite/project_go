<script lang="ts">
	import AddBookModal from "./Modal.svelte";
	import BookCard from "./BookCard.svelte";
	import Modal from "./Modal.svelte";
    let { books } = $props();
    
    let showModal = $state(false);
</script>

<div class="grid grid-cols-7">
    <button onclick={() => (showModal = true)} class="addNewBook p-4 border hover:border-2 rounded-md bg-logo-blue/20 text-8xl font-extralight hover:font-light">
        +
    </button>
    {#each books as book }
        <BookCard title={book.title} authors={book.authors} cover={book.cover}>
        </BookCard>
    {/each}
</div>
<AddBookModal bind:showModal>
    {#snippet header()}
        <h2>Add new book</h2>
    {/snippet}
    <input>
</AddBookModal>
<style>
    .addNewBook {
        width: 14rem;
        height: 20rem;
        color: var(--color-logo-blue);
        border-color: var(--color-logo-blue);
    }
</style>