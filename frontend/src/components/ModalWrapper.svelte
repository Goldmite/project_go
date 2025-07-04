<script lang="ts">
	import { scale } from 'svelte/transition';

	let { children } = $props();
</script>

<div class="relative z-30" role="dialog" tabindex="-1" aria-modal="true">
	<div class="bg-dark/20 fixed inset-0" aria-hidden="true"></div>
	<div class="fixed inset-0 z-10 w-screen">
		<div
			class="flex min-h-full items-start justify-center p-4 sm:items-center"
			in:scale={{ start: 0.23, duration: 250 }}
			out:scale={{ start: 0.77, duration: 100 }}
		>
			{@render children()}
		</div>
	</div>
</div>
