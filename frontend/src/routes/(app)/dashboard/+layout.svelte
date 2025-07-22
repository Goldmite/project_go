<script lang="ts">
	import CurrentlyReading from '../../../components/dashboard/CurrentlyReading.svelte';
	import ActivityHeatmap from '../../../components/dashboard/activity/ReadingActivity.svelte';

	let { children, data } = $props();
</script>

<div class="flex justify-center">
	<div class="mt-4 flex w-fit max-w-full flex-col gap-6">
		<ActivityHeatmap data={data.activity}></ActivityHeatmap>
		<div class="flex max-w-screen flex-wrap justify-between gap-6">
			{@render children()}
			<CurrentlyReading data={data.recent} header="Currently reading" />
		</div>
	</div>
</div>
