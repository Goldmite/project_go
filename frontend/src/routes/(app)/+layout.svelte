<script lang="ts">
	import { goto, preloadData, pushState } from '$app/navigation';
	import { page } from '$app/state';
	import AddGroupModalPage from './shared/add/+page.svelte';
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
			'bg-light dark:bg-dark hover:border-logo-red hover:shadow-logo-red h-12 w-full border p-3 font-sans font-bold shadow-lg',
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
			<div class="absolute z-10 hidden border min-w-max flex-col overflow-y-scroll h-fit max-h-60.5 sm:max-h-84.5 group-focus-within:flex group-hover:flex">
				{@render AppNavButton('+', '/shared/add')}
				{#each data.shared as shelf}
					{@render AppNavButton(shelf.name, `/shared/${shelf.name.replaceAll(' ', '%20')}`)}
				{/each}
			</div>
		</div>
	</div>
	{@render children()}
	{#if page.state.selected}
		<AddGroupModalPage data={page.state.selected} form={page.form} />
	{/if}
</div>
