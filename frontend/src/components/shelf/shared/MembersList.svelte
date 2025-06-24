<script lang="ts">
	import MemberPill from './MemberPill.svelte';

	let { members, children } = $props();
</script>

<div class="mb-6 flex">
	{#each members as m (m.id)}
		<MemberPill username={m.name} email={m.email} />
	{/each}
	{@render children()}
</div>
