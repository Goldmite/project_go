<script lang="ts">
	let { currStep } = $props();
</script>

<div class="my-4 flex justify-between">
	{#each [1, 2, 3] as step}
		<div class="relative flex flex-1 flex-col items-center text-sm">
			<div
				class="relative z-30 h-6 w-6 justify-center rounded-full border text-center transition-colors {step <=
				currStep
					? 'bg-status-logo-done'
					: 'bg-light'}"
			>
				{step}
			</div>
			{#if step < 3}
				<div class="absolute top-3 left-1/2 z-20 w-full border-b"></div>
			{/if}
		</div>
	{/each}
</div>
