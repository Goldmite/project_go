<script lang="ts">
	import { enhance } from '$app/forms';
	import type { User } from '$lib/types/user';
	import ModalWrapper from '../../../../components/ModalWrapper.svelte';
	import GroupModal from '../../../../components/shelf/shared/GroupModal.svelte';
	import InviteForm from '../../../../components/shelf/shared/InviteForm.svelte';
	import StepProgressBar from '../../../../components/StepProgressBar.svelte';
	import type { PageProps } from './$types';

	let { form }: PageProps = $props();
	let isOpen = $state(false);
	let step = $state(1);

	const invitees: User[] = $state([]);
    const invIds = new Set<string>();
	const formData = $state({
		name: '',
		emails: [] as string[]
	});
</script>

<ModalWrapper>
	<GroupModal bind:isOpen bind:step startStep=1 endStep=3 name={formData.name} emails={formData.emails} >
		<StepProgressBar currStep={step} />
		<!-- Hidden, just to catch inputs-->
		<form id="shared" method="POST" action="?/createShared" use:enhance>
			<input type="hidden" name="name" value={formData.name} />
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
				<InviteForm {form} bind:emails={formData.emails} invIds={invIds} invitees={invitees} required={formData.emails.length === 0} />
			{:else}
				<p>Ready to create!</p>
			{/if}
		</div>
	</GroupModal>
</ModalWrapper>
