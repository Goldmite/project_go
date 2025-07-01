<script lang="ts">
	let { children } = $props();
</script>

<h1 class="py-4 text-4xl font-bold hyphens-auto wrap-anywhere">
	{@render children()}
</h1>
