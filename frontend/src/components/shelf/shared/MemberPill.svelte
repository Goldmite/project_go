<script lang="ts">
	let { username, email } = $props();
</script>

<div class="group bg-logo-blue text-dark mx-1.5 my-1 h-6 rounded-full px-4 sm:mx-2.5">
	<span class="block group-active:hidden">{username}</span>
	<span class="hidden group-active:block">{email}</span>
</div>
