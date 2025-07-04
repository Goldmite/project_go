<script lang="ts">
	import '../app.css';
	import logo from '$lib/assets/SHELF_LOGO_2.svg';
	import DarkToggle from '../components/DarkToggle.svelte';
	import UserMenu from '../components/UserMenu.svelte';
	import { user } from '$lib/stores/user';

	let { children, data } = $props();
	$effect(() => {
		user.set(data.user);
	});
	let toggleUserMenu = $state(false);
	$effect(() => {
		if ($user === null) {
			toggleUserMenu = false;
		}
	});
</script>

<svelte:head>
	<script>
		const theme = localStorage.getItem('theme');
		if (theme === 'dark') {
			document.documentElement.classList.add('dark');
		}
	</script>
</svelte:head>

<nav class="flex items-center justify-between p-4">
	<div class="flex items-center space-x-4">
		<img class="h-10 w-auto" src={logo} alt="Logo" />
		<h1 class="text-3xl font-bold">Shelf</h1>
	</div>
	<div>
		<DarkToggle></DarkToggle>
		{#if $user}
			<button
				type="button"
				aria-label="User menu"
				class="rounded-full p-2 text-2xl"
				onclick={() => (toggleUserMenu = !toggleUserMenu)}
			>
				<span class="icon-[solar--user-circle-bold]"></span>
			</button>
			{#if toggleUserMenu}
				<div
					class="fixed inset-0 z-10"
					aria-hidden="true"
					onclick={() => (toggleUserMenu = false)}
				></div>
				<UserMenu userInfo={$user} />
			{/if}
		{/if}
	</div>
</nav>
{@render children()}
