<script lang="ts">
	import { goto, preloadData, pushState } from '$app/navigation';
	import { page } from '$app/state';
	import AddGroupModalPage from './shared/add/+page.svelte';
	import InviteModalPage from './shared/[id]/invite/+page.svelte';
	import type { LayoutProps } from './$types';

	let { data, children }: LayoutProps = $props();
	let currentPage = $derived(page.url.pathname);
</script>

{#snippet AppNavButton(pageName: string, pageUrl: string)}
	<button
		onclick={pageUrl != '/shared'
			? async () => {
					if (pageUrl == '/shared/add') {
						const result = await preloadData(pageUrl);
						if (result.type === 'loaded' && result.status === 200) {
							pushState(pageUrl, {
								...page.state,
								selected: result.data
							});
						}
					} else {
						goto(pageUrl);
					}
				}
			: null}
		class={[
			'bg-light dark:bg-dark hover:border-logo-red hover:shadow-logo-red min-h-12 w-full max-w-36 min-w-28 truncate border p-3 font-sans font-bold shadow-lg sm:max-w-full',
			{ 'bg-logo-red dark:bg-logo-red border-logo-red text-dark': currentPage.includes(pageUrl) }
		]}
	>
		{pageName}
	</button>
{/snippet}

<div class="m-2 sm:m-10">
	<div class="inline-flex">
		{@render AppNavButton('Dashboard', '/dashboard')}
		{@render AppNavButton('Personal', '/personal')}
		<div class="group relative top-full min-w-max">
			{@render AppNavButton('Shared', '/shared')}
			<div
				class="absolute z-10 hidden h-fit max-h-60.5 max-w-28 min-w-28 flex-col overflow-y-scroll border-y group-focus-within:flex group-hover:flex sm:max-h-84.5 sm:max-w-46"
			>
				{@render AppNavButton('+', '/shared/add')}
				{#each data.shared as shelf (shelf.id)}
					{@render AppNavButton(shelf.name, `/shared/${shelf.id}`)}
				{/each}
			</div>
		</div>
	</div>
	{@render children()}
	{#if page.state.selected}
		<AddGroupModalPage data={page.state.selected} form={page.form} />
	{/if}
	{#if page.state.invite}
		<InviteModalPage data={page.state.invite} form={page.form} />
	{/if}
</div>
