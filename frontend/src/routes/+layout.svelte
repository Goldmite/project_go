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

	let { children } = $props();
</script>

{@render children()}
