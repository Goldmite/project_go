<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/state';
	import AddGroupModal from '../../components/shelf/shared/AddGroupModal.svelte';
	import type { LayoutProps } from './$types';

	let { data, children }: LayoutProps = $props();
	let currentPage = $derived(page.url.pathname);
	let isAddGroupOpen = $state(false);
</script>

{#snippet AppNavButton(pageName: string, pageUrl: string)}
	<button
		onclick={pageUrl != '/shared' ? () => goto(pageUrl) : null}
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
			<div class="absolute z-10 hidden min-w-max flex-col group-focus-within:flex group-hover:flex">
				<button onclick={() => isAddGroupOpen = true}>
					+
				</button>
				{#each data.shared as shelf}
					{@render AppNavButton(shelf.name, `/shared/${shelf.name.replaceAll(' ', '%20')}`)}
				{/each}
			</div>
		</div>
	</div>
	{@render children()}
	{#if isAddGroupOpen}
		<AddGroupModal bind:isOpen={isAddGroupOpen}></AddGroupModal>
	{/if}
</div>
