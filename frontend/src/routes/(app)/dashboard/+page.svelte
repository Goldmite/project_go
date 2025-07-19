<script lang="ts">
	import { preloadData, replaceState } from '$app/navigation';
	import { page } from '$app/state';
	import { onMount } from 'svelte';
	import InvitationsTab from './invitations/+page.svelte';
	import StatsTab from './stats/+page.svelte';

	let selectedTab = $state('stats');
	async function preloadTab(e: any) {
		e.preventDefault();
		const { href } = e.currentTarget;
		selectedTab = href.split('/').at(-1);
		getResult(href);
	}
	async function getResult(href: string) {
		const result = await preloadData(href);
		if (result.type === 'loaded' && result.status === 200) {
			let statsData, invData;

			if (href.includes('stats')) {
				statsData = result.data;
			} else if (href.includes('invitations')) {
				invData = result.data;
			}
			replaceState('/dashboard', {
				...page.state,
				stats: statsData,
				invitations: invData
			});
		}
	}
	let invState = $state(undefined);
	$effect(() => {
		if (selectedTab === 'invitations') {
			if (page.state.invitations) invState = page.state.invitations;
		} else {
			invState = undefined;
		}
	});
	onMount(() => {
		getResult('/dashboard/stats');
	});
</script>

{#snippet tabEl(to: string)}
	<a
		class="px-4 capitalize hover:underline {selectedTab === to && 'font-semibold underline'}"
		href="/dashboard/{to}"
		data-sveltekit-preload-data="tap"
		onclick={(e) => preloadTab(e)}>{to}</a
	>
{/snippet}

<div class="w-[336px]">
	<nav class="flex justify-around p-1">
		{@render tabEl('stats')}
		{@render tabEl('invitations')}
	</nav>
	<div class="h-fit overflow-hidden rounded-2xl outline outline-current/20">
		{#if invState}
			<InvitationsTab data={invState} form={null} />
		{/if}
		{#if page.state.stats}
			<StatsTab data={page.state.stats} />
		{/if}
	</div>
</div>
