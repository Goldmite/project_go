<script lang="ts">
	import BackButton from '../../../../../components/navigation/BackButton.svelte';

	let { children } = $props();
</script>

<div class="m-4">
	<span class="absolute mt-6 text-2xl font-bold">
		<BackButton />
	</span>
	<div class="flex h-screen max-h-screen items-center justify-center">
		{@render children()}
	</div>
</div>
