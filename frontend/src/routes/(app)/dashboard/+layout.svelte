<script lang="ts">
	import CurrentlyReading from '../../../components/dashboard/CurrentlyReading.svelte';

	let { children } = $props();
</script>

{#snippet header(title: string)}
	<h1 class="text-lg font-bold sm:text-2xl">{title}</h1>
{/snippet}

<div class="flex flex-col gap-4">
	{@render children()}

	<CurrentlyReading>{@render header('Currently Reading')}</CurrentlyReading>
</div>

<style>
</style>
