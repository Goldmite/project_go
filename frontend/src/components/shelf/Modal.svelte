<script lang="ts">
	let { header = () => {}, children } = $props();
</script>

<div class="h-full w-full border-2 p-2">
	<div class="text-dark h-full">
		{@render header()}
		{@render children()}
	</div>
</div>
