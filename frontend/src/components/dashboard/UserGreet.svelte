<script lang="ts">

    let { loggedInUser } = $props();
</script>

<div class="text-xl font-serif tracking-wide italic my-4">
    {#if loggedInUser}
        Hello, {loggedInUser.name}!
    {:else}
        Loading...
    {/if}
</div>