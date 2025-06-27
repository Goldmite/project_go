<script lang="ts">
	let { children, isOpen = $bindable() } = $props();
</script>

<div class="relative z-30 {isOpen ? 'open' : ''}" role="dialog" aria-modal="true">
	<div class="bg-dark/20 fixed inset-0 transition-opacity" aria-hidden="true"></div>
	<div class="fixed inset-0 z-10 w-screen">
		<div class="flex min-h-full items-end justify-center p-4 sm:items-center">
			{@render children()}
		</div>
	</div>
</div>

<style>
	.open {
		animation: fade 0.2s ease-out;
	}
	@keyframes fade {
		from {
			opacity: 0;
		}
		to {
			opacity: 1;
		}
	}
</style>
