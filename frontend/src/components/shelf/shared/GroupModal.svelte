<script lang="ts">
	import { pushState, replaceState } from '$app/navigation';
	import { page } from '$app/state';
	import PageSubheader from '../../PageSubheader.svelte';

	let {
		children,
		step = $bindable(),
		startStep,
		endStep,
		name = '',
		invites,
		direction = $bindable()
	} = $props();

	function handleArrows(next: boolean) {
		if (next) {
			direction = 1;
			if (step == 1) {
				step = 2;
			} else if (step == 2) {
				step = 3;
			}
		} else {
			direction = -1;
			if (step == 2) {
				step = 1;
			} else if (step == 3) {
				step = 2;
			}
		}
	}
</script>

{#snippet arrowButton(next: boolean)}
	<button
		class="active:inset-shadow-md mx-4 w-full rounded-2xl p-2 text-lg hover:outline-1 focus:outline-1 active:text-base disabled:invisible {next
			? `bg-logo-blue ${step == endStep ? 'hidden' : ''}`
			: `bg-logo-red ${step == startStep ? 'hidden' : ''}`}"
		type="button"
		disabled={next
			? (step == 1 && name.length === 0) || (step == 2 && invites.length === 0)
			: false}
		form={next ? 'addShared' : null}
		onclick={() => handleArrows(next)}
	>
		{next ? '->' : '<-'}
	</button>
{/snippet}

<div
	class="bg-light text-dark flex h-[500px] w-full flex-col overflow-hidden rounded-2xl border px-8 py-4 font-sans sm:w-2/3 md:w-1/2 lg:w-2/5 2xl:w-1/4"
>
	<!-- Modal headers -->
	<div class="mx-4">
		{#if step == 1}
			<PageSubheader>I. Name your group</PageSubheader>
		{:else if step == 2}
			<PageSubheader>II. Invite your friends</PageSubheader>
		{:else}
			<PageSubheader>III. Create shared shelf</PageSubheader>
		{/if}
	</div>

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
				class="active:inset-shadow-md bg-status-logo-done mx-4 w-full rounded-2xl p-2 text-lg hover:outline-1 focus:outline-1 active:text-base disabled:invisible"
				disabled={invites.length === 0}
				type="submit"
				form="shared"
				>{'\u2713'}
			</button>
		{/if}
	</div>
</div>
