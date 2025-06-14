<script lang="ts">
    let { title, authors, cover } = $props();

</script>

<div class="card">
    <div class="imageContainer">
        <img class="h-full object-cover border rounded-md" src={cover} alt={title}>
    </div>
    <p class="font-semibold">{title}</p>
    <p class="text-sm font-light italic line-clamp-2">
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