<script lang="ts">
	import { preloadData, replaceState } from '$app/navigation';
	import { page } from '$app/state';
	import { onMount } from 'svelte';
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

{#if page.state.invitations}
	<InvitationsPage data={page.state.invitations} form={null} />
{/if}
