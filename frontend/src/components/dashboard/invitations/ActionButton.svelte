<script lang="ts">
	let { toAccept, action = $bindable('') } = $props();
</script>

<button
	class="active:inset-shadow-md size-8 rounded-full font-bold inset-shadow-sm
    outline hover:outline-2 focus:outline-2 active:text-sm disabled:pointer-events-none
    {toAccept
		? `text-status-logo-done inset-shadow-status-logo-done ${action === 'decline' && 'invisible'}`
		: `text-logo-red inset-shadow-logo-red ${action === 'accept' && 'invisible'}`}"
	type={'submit'}
	disabled={action === 'decline'}
	formaction={toAccept ? 'dashboard/invitations?/accept' : 'dashboard/invitations?/decline'}
>
	{#if toAccept}
		{#if action === 'accept'}
			{'->'}
		{:else}
			{'\u2713'}
		{/if}
	{:else}
		{'\u2717'}
	{/if}
</button>
