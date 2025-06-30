<script lang="ts">
	let { username, email, children = () => {}, hasContent } = $props();
</script>

<div class="group flex flex-row bg-logo-blue text-dark mx-1.5 my-1 h-6 rounded-full px-2 sm:mx-2.5">
	{#if !hasContent}
		<span class="block group-active:hidden">{username}</span>
		<span class="hidden group-active:block">{email}</span>
	{:else}
		<span>{email}</span>
		{@render children()}
	{/if}
</div>
