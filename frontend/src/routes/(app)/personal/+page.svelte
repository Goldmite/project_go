<script lang="ts">
	import { enhance } from '$app/forms';
	import PageHeader from '../../../components/PageHeader.svelte';
	import BookGrid from '../../../components/shelf/BookGrid.svelte';
	import AddBookModal from '../../../components/shelf/Modal.svelte';
	import type { PageProps } from './$types';

	let { data, form }: PageProps = $props();
	let isOpen = $state(false);
	let isbn = $state('');
	let addBookForm: HTMLFormElement;
	let isFormError = $derived(form?.duplicate || form?.incorrect || form?.notfound);

	function resetForm() {
		addBookForm.reset();
		isFormError = false;
	}
	$effect(() => {
		if (form?.success) {
			isOpen = false;
		}
	});
</script>

{#snippet errorMsg(msg: string)}
	<p class="text-sm/4 font-bold"><span class="bg-logo-red/90">{msg}</span></p>
{/snippet}

<PageHeader>Your shelf</PageHeader>
<BookGrid books={data.books}>
	<button
		onclick={() => (isOpen = !isOpen)}
		ontransitionend={() => resetForm()}
		class={[
			'flipCover addBookCard bg-light dark:bg-dark z-1 h-60 w-42 border p-4 text-8xl font-extralight hover:border-2 hover:font-light sm:h-84 sm:w-56',
			{ isOpen }
		]}
	>
		<div class="front bg-logo-blue/20">+</div>
		<div class="back bg-dark/30 dark:bg-light/70">x</div>
	</button>
	<div class="addBookCard bg-light absolute h-60 w-42 sm:h-84 sm:w-56">
		<AddBookModal>
			<form
				method="POST"
				bind:this={addBookForm}
				use:enhance
				class="mx-3 flex h-full flex-col space-y-6 font-serif"
			>
				<div class="w-full text-lg">
					<label for="isbn" class="block p-2">i. ISBN</label>
					<input
						class="outline-status-logo-done focus:invalid:outline-logo-red mb-1 focus:outline-3"
						bind:value={isbn}
						oninput={() => (isbn = isbn.replace(/\D/g, ''))}
						name="isbn"
						type="text"
						placeholder="Enter ISBN..."
						minlength="10"
						maxlength="13"
						required
					/>
					{#if isFormError}
						{#if form?.duplicate}
							{@render errorMsg('I. You already have this book.')}{/if}
						{#if form?.incorrect}
							{@render errorMsg('II. Must be exactly 10 or 13 digits.')}{/if}
						{#if form?.notfound}
							{@render errorMsg('III. Not found.')}{/if}
					{/if}
				</div>
				<button
					class="bg-logo-blue active:inset-shadow-md mt-auto w-full rounded-2xl p-2 text-lg hover:outline-1 focus:outline-1 active:text-base"
				>
					Add book
				</button>
			</form>
		</AddBookModal>
	</div>
</BookGrid>

<style>
	.addBookCard {
		color: var(--color-logo-blue);
		border-color: var(--color-logo-blue);
	}
	.flipCover {
		position: relative;
		transform: rotateY(0);
		transform-origin: right center;
		transition: transform 0.2s ease-out;
		transform-style: preserve-3d;
		user-select: none;
	}
	.flipCover.isOpen {
		border-width: 3px 0px 2px 3px;
		transform: rotateY(150deg) translateX(2px) skewY(-10deg);
	}
	.front,
	.back {
		display: flex;
		align-items: center;
		justify-content: center;
		left: 0;
		top: 0;
		width: 100%;
		height: 100%;
		position: absolute;
		backface-visibility: hidden;
	}
	.front {
		transform: rotateY(0);
	}
	.back {
		transform: rotateY(180deg);
	}
	.flipCover.isOpen > .front {
		transform: rotateY(0);
	}
	.flipCover.isOpen > .back {
		transform: rotateY(180deg);
	}
</style>
