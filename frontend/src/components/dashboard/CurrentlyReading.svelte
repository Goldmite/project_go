<script lang="ts">
	import { formatDateReadably, normalizeDate } from '$lib/helpers/dateWeekIndexing';
	import type { RecentlyRead } from '$lib/types/responses/book';
	import CoverWrapper from '../shelf/CoverWrapper.svelte';

	let { data, header }: { data: RecentlyRead; header: string } = $props();
	const { isbn, cover, title, authors, start_date, last_read_date } = data;
</script>

<div class="min-w-[336px]">
	<div class="p-1 text-center font-bold">{header}</div>
	{#if data}
		<div class="indented flex flex-row gap-2 outline outline-current/20">
			<a class="w-42 sm:w-56" data-sveltekit-preload-data="tap" href="/book/{isbn}/personal">
				<CoverWrapper {cover} {title} />
			</a>
			<div class="flex max-w-42 flex-col justify-between">
				<div>
					<p class="font-semibold text-wrap">{title}</p>
					<p class="line-clamp-2 text-sm font-light italic">
						{#each authors as author, i}
							{author}{i + 1 < authors.length ? ', ' : ''}
						{/each}
					</p>
				</div>
				<table class="border-t border-dashed text-current/80">
					<tbody>
						<tr>
							<td class="font-semibold">Started</td>
							<td>{formatDateReadably(normalizeDate(new Date(start_date)))}</td>
						</tr>
						<tr>
							<td class="font-semibold">Last read</td>
							<td>{formatDateReadably(normalizeDate(new Date(last_read_date)))}</td>
						</tr>
					</tbody>
				</table>
			</div>
		</div>
	{:else}
		<div class="rounded-2xl p-1.5 outline outline-current/20">
			<p class="w-full text-center text-lg tracking-wider italic">No recent book</p>
		</div>
	{/if}
</div>

<style>
</style>
