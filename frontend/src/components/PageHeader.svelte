<script lang="ts">
	let { children } = $props();
</script>

<h1 class="py-4 text-4xl font-bold wrap-anywhere hyphens-auto">
	{@render children()}
</h1>
