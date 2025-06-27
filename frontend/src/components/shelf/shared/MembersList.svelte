<script lang="ts">
	import MemberPill from './MemberPill.svelte';

	let { members, children = () => {} } = $props();
</script>

<div class="my-4 flex flex-wrap">
	{#each members as m (m.id)}
		<MemberPill username={m.name} email={m.email} />
	{/each}
	{@render children()}
</div>
