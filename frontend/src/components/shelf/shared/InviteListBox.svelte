<script lang="ts">
	import { tick } from 'svelte';
	import EyeToggle from '../../EyeToggle.svelte';
	import MembersList from './MembersList.svelte';

	let { invites = $bindable(), header, inviteToggle = $bindable(), showToggle = true } = $props();
	let div: HTMLDivElement | undefined = $state();
	let prevLength = invites.length;
	$effect.pre(() => {
		if (!div) return;
		if (invites.length <= prevLength) {
			prevLength = invites.length;
			return;
		}
		tick().then(() => {
			if (div) {
				div.scrollTo({ top: div.scrollHeight, behavior: 'smooth' });
				prevLength = invites.length;
			}
		});
	});
</script>

<div
	class="inset-shadow-md dots flex h-full flex-col overflow-hidden rounded-2xl border bg-current/7
    {invites.length !== 0 && 'alternate'}"
>
	<div
		class="bg-logo-red flex items-center justify-between border-b border-dashed px-3 brightness-90"
	>
		<h1 class="text-dark/80 font-serif text-2xl font-light tracking-tight italic text-shadow-lg/15">
			{header}
		</h1>
		{#if showToggle}
			<EyeToggle bind:toggle={inviteToggle} />
		{/if}
	</div>
	{#if invites.length !== 0}
		<div class="overflow-y-scroll" bind:this={div}>
			<MembersList bind:members={invites} isInvites={inviteToggle}></MembersList>
		</div>
	{/if}
</div>
