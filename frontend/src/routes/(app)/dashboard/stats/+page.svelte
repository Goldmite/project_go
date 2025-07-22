<script lang="ts">
	import { MINUTES_TO_SEE_STATS } from '$lib/constants.js';
	import { getPaceSpeed } from '$lib/helpers/pacePerformance';
	import type { Stats } from '$lib/types/responses/stats';
	import { slide } from 'svelte/transition';

	let { data } = $props();
	const totals: Stats = data.total;
	const hours = totals.total_time / 3600;
	const pace = totals.total_time > 0 ? Math.floor(totals.total_pages / hours) : 0;
	const time = hours >= 100 ? Math.floor(hours) : Math.floor(hours * 10) / 10;
	const performance = getPaceSpeed(pace);
</script>

<div class="flex flex-wrap gap-1.5 p-1.5 text-center" transition:slide>
	{#if totals.total_time > MINUTES_TO_SEE_STATS * 60}
		<div class="indented bg-logo-purple/20 grow">
			<h1 class="font-serif text-4xl text-shadow-lg">{totals.total_pages}</h1>
			<span class="text-sm italic">total pages read</span>
		</div>
		<div class="inset-shadow-md hidden flex-1 rounded-2xl"></div>
		<div class="flex grow gap-1.5">
			<div class="indented bg-logo-purple/20 grow">
				<h1 class="font-serif text-4xl text-shadow-lg">{time}</h1>
				<span class="text-sm italic">total hours read</span>
			</div>
			<div
				class="indented shrink {performance === 'fast'
					? 'bg-status-logo-done/20'
					: performance === 'average'
						? 'bg-status-logo-progress/20'
						: 'bg-logo-red/20'}"
			>
				<h1 class="font-serif text-4xl text-shadow-lg">{pace}</h1>
				<span class="text-sm italic">pages/h</span>
			</div>
		</div>
	{:else}
		<p class="w-full text-center text-lg tracking-wider italic">No data</p>
	{/if}
</div>
