<script lang="ts">
	import { enhance } from '$app/forms';
	import { getInviteState, setInviteState } from '$lib/inviteState.svelte';
	import ModalWrapper from '../../../../components/ModalWrapper.svelte';
	import GroupModal from '../../../../components/shelf/shared/GroupModal.svelte';
	import InviteForm from '../../../../components/shelf/shared/InviteForm.svelte';
	import StepProgressBar from '../../../../components/StepProgressBar.svelte';
	import type { PageProps } from './$types';

	let { form }: PageProps = $props();
	let step = $state(1);
	let isOpen = $state(false);
	setInviteState();
	const inviteState = getInviteState();
	let name = $state('');
</script>

<ModalWrapper>
	<GroupModal bind:isOpen bind:step startStep="1" endStep="3" {name} invites={inviteState.invites}>
		<StepProgressBar currStep={step} />
		<!-- Hidden, just to catch inputs-->
		<form id="shared" method="POST" action="?/createShared" use:enhance>
			<input type="hidden" name="name" value={name} />
			{#each inviteState.invites as invite}
				<input type="hidden" name="emails[]" value={invite.email} />
			{/each}
		</form>

		<div class="mx-4 mt-4 flex flex-1 flex-col">
			{#if step == 1}
				<input
					class="outline-status-logo-done focus:invalid:outline-logo-red focus:outline-3"
					bind:value={name}
					type="text"
					placeholder="Be creative..."
					maxlength="20"
					required
				/>
			{:else if step == 2}
				<InviteForm {form} required={inviteState.invites.length === 0} />
			{:else}
				<p>Ready to create!</p>
			{/if}
		</div>
	</GroupModal>
</ModalWrapper>
