<script lang="ts">
	import { preloadData, replaceState } from '$app/navigation';
	import { page } from '$app/state';
	import { onMount } from 'svelte';
	import CurrentlyReading from '../../../components/dashboard/CurrentlyReading.svelte';
	import InvitationsPage from './invitations/+page.svelte';

	onMount(async () => {
		const result = await preloadData('/dashboard/invitations');
		if (result.type === 'loaded' && result.status === 200) {
			replaceState('/dashboard/invitations', {
				...page.state,
				invitations: result.data
			});
		}
	});
</script>

{#snippet header(title: string)}
	<h1 class="text-lg font-bold sm:text-2xl">{title}</h1>
{/snippet}

<div class="flex justify-between">
	<div class="flex-1">
		{@render header('Shelves')}
	</div>
	{#if page.state.invitations}
		<div class="h-fit overflow-hidden rounded-2xl outline">
			<div class="-mr-0.5 h-fit max-h-50 min-h-10 w-[336px] overflow-y-scroll">
				<InvitationsPage data={page.state.invitations} form={null} />
			</div>
		</div>
	{/if}
	<CurrentlyReading>{@render header('Currently Reading')}</CurrentlyReading>
</div>

<style>
</style>
