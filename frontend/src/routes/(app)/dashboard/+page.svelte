<script lang="ts">
	import { preloadData, pushState } from '$app/navigation';
	import { page } from '$app/state';
	import CurrentlyReading from '../../../components/dashboard/CurrentlyReading.svelte';
    import InvitationsPage from '../dashboard/invitations/+page.svelte';
</script>

{#snippet header(title: string)}
	<h1 class="text-lg font-bold sm:text-2xl">{title}</h1>
{/snippet}

<div class="flex justify-between">
	<div class="flex-1">
		{@render header('Shelves')}
		<button
            onclick={async () => {
                const result = await preloadData('dashboard/invitations');
                if (result.type === 'loaded' && result.status === 200) {
                    pushState('dashboard/invitations', {
                        ...page.state,
                        invitations: result.data
                    });
                }
            }}    
        >open</button>
	</div>
    {#if page.state.invitations}
        <div class="border rounded-2xl overflow-hidden h-[400px]">
            <InvitationsPage data={page.state.invitations} form={null} />
        </div>
    {/if}
	<CurrentlyReading>{@render header('Currently Reading')}</CurrentlyReading>
</div>

<style>
</style>
