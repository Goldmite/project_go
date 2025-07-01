<script lang="ts">
	import { applyAction, enhance } from '$app/forms';
	import { getInviteState } from '$lib/inviteState.svelte';
	import ErrorMsg from '../../errors/ErrorMsg.svelte';
	import InviteListBox from './InviteListBox.svelte';

	const MAX_INVITES = 7;
	let { invited = [], members = [], form, required } = $props();
	const inviteState = getInviteState();

	let inviteToggle = $state(true);
	let [isDuplicate, isMember] = $derived.by((): [boolean, boolean] => {
		if (form?.invitee !== undefined) {
			return [ 
				invited.some((user) => user.id === form.invitee.id), 
				members.some((m) => m.id === form.invitee.id) ];
		}
		return [ false, false ];
	});
	let canSubmit = $state(false);
	let input = $state('');
	$effect(() => {
		if (
			form?.fakesuccess &&
			canSubmit &&
			!inviteState.isExist(form?.invitee.id) &&
			inviteState.invites.length < MAX_INVITES &&
			!isDuplicate && !isMember
		) {
			input = '';
			inviteState.add(form.invitee.id, form.invitee.name, form.invitee.email);
			canSubmit = false;
		}
	});
</script>

<form
	method="POST"
	action="?/checkUser"
	use:enhance={() => {
		return async ({ result }) => {
			if (result.type === 'failure') {
				canSubmit = true;
			}
			await applyAction(result);
		};
	}}
>
	<div class="flex w-full flex-row">
		<input
			class="outline-status-logo-done focus:invalid:outline-logo-red w-4/5 focus:outline-3"
			name="email"
			type="email"
			bind:value={input}
			placeholder="Add friends..."
			maxlength="256"
			{required}
		/>
		<button
			class="bg-status-logo-done active:inset-shadow-md ml-2 w-1/7 rounded-xl text-2xl hover:outline-1 focus:outline-1 active:text-xl"
			>+</button
		>
	</div>
</form>

<div class="my-2 h-8 flex-none">
	{#if form?.notfound}
		<ErrorMsg msg="I. User not found."></ErrorMsg>
	{/if}
	{#if isDuplicate}
		<ErrorMsg msg="II. Already invited."></ErrorMsg>
	{/if}
	{#if form?.yourself}
		<ErrorMsg msg="III. No need to invite yourself."></ErrorMsg>
	{/if}
	{#if isMember}
		<ErrorMsg msg="IV. Already a member."></ErrorMsg>
	{/if}
	{#if inviteState.invites.length === MAX_INVITES && input != ''}
		<ErrorMsg msg="V. Max invites reached."></ErrorMsg>
	{/if}
</div>

<div class="h-43">
	{#if inviteToggle}
		<InviteListBox
			bind:invites={inviteState.invites}
			header="Will invite"
			bind:inviteToggle
			showToggle={invited.length !== 0}
		></InviteListBox>
	{:else}
		<InviteListBox bind:invites={invited} header="Already invited" bind:inviteToggle
		></InviteListBox>
	{/if}
</div>
