<script lang="ts">
	let { children } = $props();
</script>

<h1 class="p-4 text-4xl font-bold">
	{@render children()}
</h1>
