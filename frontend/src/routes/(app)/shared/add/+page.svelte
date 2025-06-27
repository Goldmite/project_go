<script lang="ts">
	import { enhance } from '$app/forms';
	import type { User } from '$lib/types/user';
	import ErrorMsg from '../../../../components/errors/ErrorMsg.svelte';
	import ModalWrapper from '../../../../components/ModalWrapper.svelte';
	import PageSubheader from '../../../../components/PageSubheader.svelte';
	import MembersList from '../../../../components/shelf/shared/MembersList.svelte';
	import StepProgressBar from '../../../../components/StepProgressBar.svelte';
	import type { PageProps } from './$types';

	const maxInvites = 3;
	let { form }: PageProps = $props();
	let isOpen = $state(false);
	let step = $state(1);

	let emailInput = $state('');
	const invitees: User[] = $state([]);
	const invIds = new Set<string>();
	
	const formData = $state({
        name: '',
        emails: [] as string[]
    });

	function handleArrows(next: boolean) {
		if (next) {
			if (step == 1) {
				step = 2;
			} else if (step == 2) {
				step = 3;
			}
		} else {
			if (step == 2) {
				step = 1;
			} else if (step == 3) {
				step = 2;
			}
		}
	}
	
	$effect(() => {
		if (form?.fakesuccess && !invIds.has(form.invitee.id) && invitees.length < maxInvites) {
			emailInput = '';
			invitees.push({
				id: form.invitee.id,
				name: form.invitee.name,
				email: form.invitee.email
			});
			formData.emails = [...formData.emails, form.invitee.email];
			invIds.add(form.invitee.id);
		}
	});
</script>

{#snippet arrowButton(next: boolean)}
	<button
		class="active:inset-shadow-md mx-4 w-full rounded-2xl p-2 text-lg hover:outline-1 focus:outline-1 active:text-base disabled:invisible {next
			? `bg-logo-blue ${step == 3 ? 'hidden' : ''}`
			: `bg-logo-red ${step == 1 ? 'hidden' : ''}`}"
		disabled={next ? (step == 1 && formData.name.length === 0) || (step == 2 && formData.emails.length === 0) : false}
		type="button" form={next ? "addShared" : null}
		onclick={() => handleArrows(next)}
	>
		{next ? '->' : '<-'}
	</button>
{/snippet}

<ModalWrapper>
	<div
		class="bg-light text-dark flex h-[500px] w-full flex-col rounded-2xl border px-8 py-4 font-sans sm:w-2/3 md:w-1/2 lg:w-2/5 2xl:w-1/4"
	>
		<StepProgressBar currStep={step} />
		<div class="mx-4">
			{#if step == 1}
				<PageSubheader>I. Name your group</PageSubheader>
			{:else if step == 2}
				<PageSubheader>II. Invite your friends</PageSubheader>
			{:else}
				<PageSubheader>III. Create shared shelf</PageSubheader>
			{/if}
		</div>
		<!-- Hidden, just to catch inputs-->
		<form id="addShared" method="POST" action="?/createShared" use:enhance >
			<input type="hidden" name="name" value={formData.name} >
			{#each formData.emails as email}
				<input type="hidden" name="emails[]" value={email} />
			{/each}
		</form>

		<div class="m-4 flex flex-1 flex-col">
			{#if step == 1}
			<input
				class="outline-status-logo-done focus:invalid:outline-logo-red focus:outline-3"
				bind:value={formData.name}
				type="text"
				placeholder="Be creative..."
				maxlength="20"
				required
			/>
			{:else if step == 2}
			<form method="POST" action="?/checkUser" use:enhance>
			<div class="flex w-full flex-row">
				<input
					class="outline-status-logo-done focus:invalid:outline-logo-red w-4/5 focus:outline-3"
					name="email"
					type="email"
					bind:value={emailInput}
					placeholder="Add friends..."
					maxlength="256"
					required={formData.emails.length === 0}
				/>
				<button class="ml-2 w-1/7 text-2xl bg-status-logo-done active:inset-shadow-md rounded-xl hover:outline-1 focus:outline-1 active:text-xl">+</button>
			</div>
			</form>
			<div class="flex-none h-8 my-2">
			{#if form?.notfound}
				<ErrorMsg msg="I. User not found."></ErrorMsg>
			{/if}
			{#if form?.duplicate}
				<ErrorMsg msg="II. No need to invite yourself."></ErrorMsg>
			{/if}
			{#if invitees.length === maxInvites && emailInput != ''}
				<ErrorMsg msg="III. Max invites reached."></ErrorMsg>
			{/if}
			</div>
			<div class="flex-1 border rounded-2xl inset-shadow-md backdrop-brightness-95 dots {invitees.length !== 0 && 'alternate'}">
			{#if invitees.length !== 0}
				<MembersList members={invitees}></MembersList>
			{/if}
			</div>
			{:else}
                <p>Ready to create!</p>
			{/if}
		</div>

		<div class="mt-auto flex w-full justify-around">
			{#if step == 1}
			<button
				class="active:inset-shadow-md bg-logo-red mx-4 w-full rounded-2xl p-2 text-lg hover:outline-1 focus:outline-1 active:text-base"
				onclick={() => history.back()}>{'\u2717'}</button
			>
			{/if}
			{@render arrowButton(false)}
			{@render arrowButton(true)}
			{#if step == 3}
			<button
				class="active:inset-shadow-md bg-logo-blue mx-4 w-full rounded-2xl p-2 text-lg hover:outline-1 focus:outline-1 active:text-base"
				onclick={() => (isOpen = false)}
				type="submit"
				form="addShared"
				>{'\u2713'} </button
			>
			{/if}
		</div>
	</div>
</ModalWrapper>
