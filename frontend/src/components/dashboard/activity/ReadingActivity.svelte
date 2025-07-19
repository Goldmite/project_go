<script lang="ts">
	import { mapToWeekIndexing } from '$lib/helpers/dateWeekIndexing';
	import { formatMinToHoursOrMin } from '$lib/helpers/formatMinTime';
	import type { ReadingSession } from '$lib/types/responses/stats';
	import { onMount } from 'svelte';
	import MonthLabel from './MonthLabel.svelte';

	let { data }: { data: ReadingSession[] } = $props();
	const gridData = data ? mapToWeekIndexing(data) : {};
	let usedMonths: string[] = [];

	let scrollable: HTMLDivElement | undefined;
	onMount(() => {
		if (scrollable) scrollable.scrollTo(scrollable.scrollWidth, 0);
	});
</script>

<div
	bind:this={scrollable}
	class="indented max-w-screen overflow-x-scroll outline outline-current/20 md:max-w-[860px]"
>
	<div class="grid w-[844px] grid-cols-[repeat(auto-fit,0.75rem)] gap-1">
		{#each Object.keys(gridData) as week}
			<div class="grid w-3 grid-rows-[repeat(8,0.75rem)] gap-1">
				<MonthLabel gridCol={gridData[+week]} {usedMonths}></MonthLabel>
				{#each Object.entries(gridData[+week]) as [, cell]}
					<div
						id={cell.date}
						tabindex="-1"
						data-tooltip="{formatMinToHoursOrMin(cell.time / 60)} on {cell.date.replace(
							/ \d{2}(\d{2})/,
							"'$1"
						)}"
						class="cell
            {+week < 3 ? 'after:left-0' : +week > 49 ? 'after:right-0' : 'after:-translate-x-1/2'} 
                {cell.level == 1
							? 'bg-activity-lighter dark:bg-activity-darker'
							: cell.level == 2
								? 'bg-activity-light dark:bg-activity-dark'
								: cell.level == 3
									? 'bg-activity-mid'
									: cell.level == 4
										? 'bg-activity-dark dark:bg-activity-light'
										: cell.level == 5
											? 'bg-activity-darker dark:bg-activity-lighter'
											: 'bg-current/7'}"
					></div>
				{/each}
			</div>
		{/each}
	</div>
	<div class="flex flex-row items-center gap-1.5 pt-2">
		<span class="text-xs text-current/60">{'<30m'}</span>
		<div class="legend bg-activity-lighter dark:bg-activity-darker"></div>
		<div class="legend bg-activity-light dark:bg-activity-dark"></div>
		<div class="legend bg-activity-mid"></div>
		<div class="legend bg-activity-dark dark:bg-activity-light"></div>
		<div class="legend bg-activity-darker dark:bg-activity-lighter"></div>
		<span class="text-xs text-current/60">{'>2h'}</span>
	</div>
</div>

<style>
	.legend {
		width: 14px;
		height: 14px;
		border-radius: 2px;
	}
	.cell {
		width: 12px;
		height: 12px;
		border-radius: 2px;
		position: relative;
		border-color: transparent;
		transition:
			border-color 0.5s ease 0.3s,
			border-width 0.5s ease 0.3s;
	}
	.cell::after {
		content: attr(data-tooltip);
		position: absolute;
		background-color: color-mix(in oklab, currentcolor 20%, var(--color-dark));
		color: var(--color-light);
		border: 1px solid color-mix(in oklab, currentcolor 50%, transparent);
		top: -29px;
		padding: 1px;
		font-size: 12px;
		border-radius: 4px;
		white-space: nowrap;
		display: none;
		pointer-events: none;
		z-index: 20;
	}
	:root.dark .cell::after {
		background-color: color-mix(in oklab, currentcolor 20%, var(--color-light));
		color: var(--color-dark);
	}

	.cell:hover::after {
		display: block;
	}
	.cell:focus::after {
		display: block;
	}
	.cell:hover,
	.cell:focus {
		transition: border-color 0s;
		border: 6px solid color-mix(in oklab, currentcolor 20%, transparent);
	}
	.cell:focus {
		z-index: 10;
	}
</style>
