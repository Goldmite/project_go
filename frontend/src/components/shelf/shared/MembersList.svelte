<script lang="ts">
	import MemberPill from "./MemberPill.svelte";

    let { members, children } = $props();

</script>

<div class="flex mb-6">
    {#each members as m (m.id)}
        <MemberPill username={m.name} email={m.email} />
    {/each}
    {@render children()}
</div>