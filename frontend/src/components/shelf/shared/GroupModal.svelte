<script lang="ts">
	let { children, isOpen = $bindable(), step = $bindable(), startStep, endStep } = $props();

	function handleArrows(next: boolean) {
		if (next) {
			if (step == 1) {
				step = 2;
			} else if (step == 2) {
				step = 3;
			}
		} else {
			if (step == 2) {
				step = 1;
			} else if (step == 3) {
				step = 2;
			}
		}
	}
</script>

{#snippet arrowButton(next: boolean)}
	<!-- disabled={next ? (step == 1 && formData.name.length === 0) || (step == 2 && formData.emails.length === 0) : false} -->
	<button
		class="active:inset-shadow-md mx-4 w-full rounded-2xl p-2 text-lg hover:outline-1 focus:outline-1 active:text-base disabled:invisible {next
			? `bg-logo-blue ${step == endStep ? 'hidden' : ''}`
			: `bg-logo-red ${step == startStep ? 'hidden' : ''}`}"
		type="button"
		form={next ? 'addShared' : null}
		onclick={() => handleArrows(next)}
	>
		{next ? '->' : '<-'}
	</button>
{/snippet}

<div
	class="bg-light text-dark flex h-[500px] w-full flex-col rounded-2xl border px-8 py-4 font-sans sm:w-2/3 md:w-1/2 lg:w-2/5 2xl:w-1/4"
>
	<!-- Modal content -->
	{@render children()}
	<!-- Modal navigation buttons -->
	<div class="mt-auto flex w-full justify-around">
		{#if step == startStep}
			<button
				class="active:inset-shadow-md bg-logo-red mx-4 w-full rounded-2xl p-2 text-lg hover:outline-1 focus:outline-1 active:text-base"
				onclick={() => history.back()}>{'\u2717'}</button
			>
		{/if}
		{@render arrowButton(false)}
		{@render arrowButton(true)}
		{#if step == endStep}
			<button
				class="active:inset-shadow-md bg-logo-blue mx-4 w-full rounded-2xl p-2 text-lg hover:outline-1 focus:outline-1 active:text-base"
				onclick={() => (isOpen = false)}
				type="submit"
				form="addShared"
				>{'\u2713'}
			</button>
		{/if}
	</div>
</div>
