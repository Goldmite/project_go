<script lang="ts">
	import { enhance } from "$app/forms";
	import type { User } from "$lib/types/user";
	import ModalWrapper from "../../../../../components/ModalWrapper.svelte";
	import GroupModal from "../../../../../components/shelf/shared/GroupModal.svelte";
	import InviteForm from "../../../../../components/shelf/shared/InviteForm.svelte";
	import type { PageProps } from "./$types";

    let { form }: PageProps = $props();
    let isOpen = $state(false);
    const invitees: User[] = $state([]);
    const invIds = new Set<string>();
    let emails: string[] = $state([]);
    
</script>

<ModalWrapper>
    <GroupModal bind:isOpen step=2 startStep=2 endStep=2 emails={emails}>
        <!-- Hidden, just to catch inputs-->
		<form id="shared" method="POST" action="?/invite" use:enhance>
			{#each emails as email}
				<input type="hidden" name="emails[]" value={email} />
			{/each}
		</form>
		<div class="m-4 flex flex-1 flex-col">
            <InviteForm {form} bind:emails={emails} invIds={invIds} invitees={invitees} required={emails.length === 0} />
        </div>
    </GroupModal>
</ModalWrapper>