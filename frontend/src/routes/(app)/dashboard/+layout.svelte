<script lang="ts">
	import { getUserState } from '$lib/userState.svelte';
	import CurrentlyReading from '../../../components/dashboard/CurrentlyReading.svelte';
	import UserGreet from '../../../components/dashboard/UserGreet.svelte';

	let { children } = $props();
	const loggedInUser = getUserState().user;
</script>

{#snippet header(title: string)}
	<h1 class="text-lg font-bold sm:text-2xl">{title}</h1>
{/snippet}

<div class="flex flex-col gap-4">
	<UserGreet {loggedInUser} />
	{@render children()}
	
	<CurrentlyReading>{@render header('Currently Reading')}</CurrentlyReading>
</div>

<style>
</style>
