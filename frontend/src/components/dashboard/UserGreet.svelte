<script lang="ts">
	let { loggedInUser } = $props();
</script>

<div class="my-4 font-serif text-xl tracking-wide italic">
	{#if loggedInUser}
		Hello, {loggedInUser.name}!
	{:else}
		Loading...
	{/if}
</div>
