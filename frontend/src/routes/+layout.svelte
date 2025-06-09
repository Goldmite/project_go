<svelte:head>
	<script>
		// Apply saved theme before app renders
		const theme = localStorage.getItem('theme');
		if (theme === 'dark') {
		document.documentElement.classList.add('dark');
		}
</script>
</svelte:head>

<script lang="ts">
	import '../app.css';
	import DarkToggle from "../components/DarkToggle.svelte";

	let { children } = $props();
</script>

<nav class="flex p-4 items-center justify-between">
    <div class="flex items-center space-x-4">
        <img class="h-10 w-auto" src="./SHELF_LOGO.svg" alt="Logo" >
        <h1 class="text-3xl font-bold">Shelf</h1>
    </div>
    <DarkToggle> </DarkToggle>
</nav>
{@render children()}
