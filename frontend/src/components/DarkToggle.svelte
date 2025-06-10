<script>
	import { onMount } from 'svelte';

	let dark = true;

	onMount(() => {
		dark = localStorage.getItem('theme') === 'dark';
		updateTheme();
	});

	function toggleTheme() {
		dark = !dark;
		localStorage.setItem('theme', dark ? 'dark' : 'light');
		updateTheme();
	}

	function updateTheme() {
		const root = document.documentElement;
		root.classList.toggle('dark', dark);
	}
</script>

<button
	type="button"
	aria-label="Toggle dark mode"
	class="rounded-full p-2 text-2xl"
	onclick={toggleTheme}
>
	{#if dark}
		<span class="icon-[solar--sun-2-bold]"></span>
	{:else}
		<span class="icon-[solar--moon-stars-bold]"></span>
	{/if}
</button>
