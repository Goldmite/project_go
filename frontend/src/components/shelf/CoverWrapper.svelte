<script lang="ts">
	let { cover, title } = $props();
</script>

<div class="h-60 sm:h-84">
	<img class="h-full rounded-md border border-current/20 object-cover" src={cover} alt={title} />
</div>
