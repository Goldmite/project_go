<script lang="ts">
	import { enhance } from '$app/forms';
	import ErrorMsg from '../../errors/ErrorMsg.svelte';
	import MembersList from './MembersList.svelte';

	const MAX_INVITES = 7;
	let { form, emails = $bindable(), required, invitees, invIds } = $props();

	let input = $state('');
	$effect(() => {
		if (form?.fakesuccess && !invIds.has(form.invitee.id) && invitees.length < MAX_INVITES) {
			input = '';
			invitees.push({
				id: form.invitee.id,
				name: form.invitee.name,
				email: form.invitee.email
			});
            //emails.push(form.invitee.email);
			emails = [...emails, form.invitee.email];
			invIds.add(form.invitee.id);
		}
	});
</script>

<form method="POST" action="?/checkUser" use:enhance>
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
	{#if form?.duplicate}
		<ErrorMsg msg="II. No need to invite yourself."></ErrorMsg>
	{/if}
	{#if invitees.length === MAX_INVITES && input != ''}
		<ErrorMsg msg="III. Max invites reached."></ErrorMsg>
	{/if}
</div>

<div
	class="inset-shadow-md dots flex-1 rounded-2xl border backdrop-brightness-95
    {invitees.length !== 0 && 'alternate'}"
>
	{#if invitees.length !== 0}
		<MembersList members={invitees}></MembersList>
	{/if}
</div>
