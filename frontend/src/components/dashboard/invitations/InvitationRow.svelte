<script>
	import { applyAction, enhance } from '$app/forms';
	import { toRoman } from '$lib/helpers/romanNumerals';
	import { redirect } from '@sveltejs/kit';
	import ActionButton from './ActionButton.svelte';

	let { inv, row } = $props();
	let action = $state('');
	let isDecided = $state(false);
</script>

<tr
	class="group border border-current/40 {!isDecided
		? ''
		: action === 'accept'
			? 'bg-status-logo-done/20'
			: action === 'decline'
				? 'bg-logo-red/20'
				: 'bg-current/10'}"
>
	<td class="pl-1 text-sm font-light group-hover:bg-current/5">{toRoman(row)}.</td>
	<td class=" border-r border-current/10 px-1 group-hover:bg-current/5">
		{inv.inviter_name}
		<span class="text-xs tracking-tighter text-current/50 italic">invited to</span>
		<span class="">{inv.group_name}</span>
	</td>
	<td class="p-1 pr-1.5 group-hover:bg-current/5">
		<form
			method="POST"
			class="flex flex-row justify-center gap-3"
			use:enhance={() => {
				return async ({ result }) => {
					if (result.type === 'success') {
						isDecided = true;
						if (result.data?.isAccepted) action = 'accept';
						else action = 'decline';
					}
					await applyAction(result);
				};
			}}
		>
			<input type="hidden" name="group_id" value={inv.group_id} />
			<input type="hidden" name="redirect" value={isDecided} />
			<ActionButton toAccept={false} bind:action />
			<ActionButton toAccept={true} bind:action />
		</form>
	</td>
</tr>
