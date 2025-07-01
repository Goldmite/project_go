<script lang="ts">
	import { getInviteState } from '$lib/inviteState.svelte';
	import type { User } from '$lib/types/user';
	import MemberPill from './MemberPill.svelte';

	let { members = $bindable(), isInvites = false } = $props();
	const inviteState = getInviteState();

	function handleRemove(id: string) {
		inviteState.remove(id);
		members = members.filter((m: User) => m.id !== id);
	}
</script>

<div class="my-2 flex flex-wrap">
	{#each members as member (member.id)}
		<MemberPill username={member.name} email={member.email} hasContent={isInvites}>
			{#if isInvites}
				<button onclick={() => handleRemove(member.id)}>
					{'\u2717'}
				</button>
			{/if}
		</MemberPill>
	{/each}
</div>
