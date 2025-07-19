<script lang="ts">
	import CurrentlyReading from '../../../components/dashboard/CurrentlyReading.svelte';
	import ActivityHeatmap from '../../../components/dashboard/activity/ReadingActivity.svelte';

	let { children, data } = $props();
</script>

{#snippet header(title: string)}
	<h1 class="text-lg font-bold sm:text-2xl">{title}</h1>
{/snippet}

<div class="mt-4 flex flex-col gap-4">
	<ActivityHeatmap data={data.activity}></ActivityHeatmap>

	{@render children()}

	<CurrentlyReading>{@render header('Currently Reading')}</CurrentlyReading>
</div>
