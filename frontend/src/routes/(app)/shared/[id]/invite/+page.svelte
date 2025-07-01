<script lang="ts">
	import { enhance } from '$app/forms';
	import { page } from '$app/state';
	import { getInviteState, setInviteState } from '$lib/inviteState.svelte';
	import ModalWrapper from '../../../../../components/ModalWrapper.svelte';
	import GroupModal from '../../../../../components/shelf/shared/GroupModal.svelte';
	import InviteForm from '../../../../../components/shelf/shared/InviteForm.svelte';
	import type { PageProps } from './$types';

	let { data, form }: PageProps = $props();
	let isOpen = $state(false);
	setInviteState();
	const inviteState = getInviteState();
</script>

<ModalWrapper>
	<GroupModal bind:isOpen step="2" startStep="2" endStep="2" invites={inviteState.invites}>
		<!-- Hidden, just to catch inputs-->
		<form id="shared" method="POST" action="?/invite" use:enhance>
			{#each inviteState.invites as invite}
				<input type="hidden" name="emails[]" value={invite.email} />
			{/each}
		</form>
		<div class="mx-4 mt-4 flex flex-1 flex-col">
			<InviteForm invited={data.invitedUsers} members={page.state.groupMembers} {form} required={inviteState.invites.length === 0} />
		</div>
	</GroupModal>
</ModalWrapper>
