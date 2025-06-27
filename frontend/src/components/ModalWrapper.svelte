<script lang="ts">
	let { children } = $props();
</script>

<div class="relative z-30" role="dialog" aria-modal="true">
	<div class="bg-dark/20 fixed inset-0" aria-hidden="true"></div>
	<div class="fixed inset-0 z-10 w-screen">
		<div class="flex min-h-full items-start justify-center p-4 sm:items-center">
			{@render children()}
		</div>
	</div>
</div>
