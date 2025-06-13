<script lang="ts">
	import BookCard from "./BookCard.svelte";
    let { books } = $props();
    /*const isbn = "1234567890123"
    const title = "Test title is quite long and will wrap nicely";
    const authors = ["First Author name", "Second Author longer name", "Third author name not fit"];
    const cover = "";//"http://books.google.com/books/content?id=SzF8AwAAQBAJ&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api";
    const b = {isbn, title ,authors, cover}
    const books = [
        b, b, b, b, b, b, b, b
    ]*/
</script>

<div class="grid grid-cols-7">
    {#each books as book }
        <BookCard title={book.title} authors={book.authors} cover={book.cover}>
        </BookCard>
    {:else}
        <p>You have no books on your shelf.</p>
    {/each}
</div>
