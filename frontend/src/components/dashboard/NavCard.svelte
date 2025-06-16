<script lang="ts">
    let { children } = $props();

</script>

<button class="card w-42 h-42 sm:w-56 sm:h-56">
    {@render children()}
</button>

<style>
    .card {
        border-radius: var(--radius-2xl);
        border-width: 1px;
        box-shadow: inset 2px 10px 6px;
    }
</style>