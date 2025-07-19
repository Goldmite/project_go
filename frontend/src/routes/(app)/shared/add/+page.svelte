<script lang="ts">
	import { applyAction, enhance } from '$app/forms';
	import { getInviteState, setInviteState } from '$lib/inviteState.svelte';
	import { fly } from 'svelte/transition';
	import ModalWrapper from '../../../../components/ModalWrapper.svelte';
	import GroupModal from '../../../../components/shelf/shared/GroupModal.svelte';
	import InviteForm from '../../../../components/shelf/shared/InviteForm.svelte';
	import StepProgressBar from '../../../../components/StepProgressBar.svelte';
	import type { PageProps } from './$types';
	import { goto, invalidate } from '$app/navigation';

	let { form }: PageProps = $props();
	let step = $state(1);
	setInviteState();
	const inviteState = getInviteState();
	let name = $state('');
	let middleTransitionDirection = $state(1); //1 forward, -1 backward
</script>

<ModalWrapper>
	<GroupModal
		bind:step
		startStep="1"
		endStep="3"
		{name}
		invites={inviteState.invites}
		bind:direction={middleTransitionDirection}
	>
		<StepProgressBar currStep={step} />
		<!-- Hidden, just to catch inputs-->
		<form
			id="shared"
			method="POST"
			action="?/createShared"
			use:enhance={() => {
				return async ({ result }) => {
					if (result.type === 'success') {
						await invalidate('app:newgroup');
						goto(`/shared/${result.data?.groupId}`);
					}
					await applyAction(result);
				};
			}}
		>
			<input type="hidden" name="name" value={name} />
			{#each inviteState.invites as invite}
				<input type="hidden" name="emails[]" value={invite.email} />
			{/each}
		</form>

		<div class="mx-4 mt-4 flex flex-1 flex-col">
			{#if step == 1}
				<input
					class="outline-status-logo-done focus:invalid:outline-logo-red focus:outline-3"
					in:fly={{ delay: 150, duration: 150, x: -300 }}
					out:fly={{ delay: 0, duration: 150, x: -300 }}
					bind:value={name}
					type="text"
					placeholder="Be creative..."
					maxlength="20"
					required
				/>
			{:else if step == 2}
				<div
					in:fly={{ delay: 150, duration: 150, x: 300 * middleTransitionDirection }}
					out:fly={{ delay: 0, duration: 150, x: -300 * middleTransitionDirection }}
				>
					<InviteForm {form} required={inviteState.invites.length === 0} />
				</div>
			{:else}
				<div
					in:fly={{ delay: 150, duration: 150, x: 300 }}
					out:fly={{ delay: 0, duration: 150, x: 300 }}
				>
					<p>Ready to create!</p>
				</div>
			{/if}
		</div>
	</GroupModal>
</ModalWrapper>
