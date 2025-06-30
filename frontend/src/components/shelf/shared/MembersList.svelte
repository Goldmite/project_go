<script lang="ts">
	import type { User } from '$lib/types/user';
	import MemberPill from './MemberPill.svelte';

	let { members, children = () => {}, inviteState = undefined } = $props();

	function handleRemove(id: string) {
		inviteState.remove(id);
		members = members.filter((m: User) => m.id !== id);
	}
</script>

<div class="my-4 flex flex-wrap">
	{@render children()}
	{#each members as member (member.id)}
		<MemberPill username={member.name} email={member.email} hasContent={inviteState != undefined}>
			{#if inviteState}
				<button onclick={() => handleRemove(member.id)}>
					{'\u2717'}
				</button>
			{/if}
		</MemberPill>
	{/each}
</div>
