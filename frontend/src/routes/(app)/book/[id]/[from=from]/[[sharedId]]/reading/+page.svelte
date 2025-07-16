<script lang="ts">
	import { slide } from 'svelte/transition';
	import type { Snapshot } from './$types';
	import { Timer } from '$lib/timerState.svelte';
	import { enhance } from '$app/forms';
	import type { BookProgress } from '$lib/types/responses/book';
	import PageSubheader from '../../../../../../../components/PageSubheader.svelte';
	import { getPaceSpeed } from '$lib/helpers/pacePerformance';

	const timer = new Timer();
	let ts = $derived(timer.state);
	export const snapshot: Snapshot<typeof timer.state> = {
		capture: () => timer.toJSON(),
		restore: (state) => {
			timer.state = Timer.fromJSON(state);
		}
	};
	// Current progress data
	let { data } = $props();
	const progress: BookProgress = data.progress;
	const actualStartPage = progress.current_page === 0 ? progress.first_page : progress.current_page;
	const actualPace = progress.pages_read / (progress.time_read / 60);
	let estimateEndPage = $state(actualStartPage);
	// Time input in case of edit and validation
	let submitTime = $state(0);
	let editTime = $state(false);
	let timeMinsInput = $state('');
	function checkTimeInput() {
		editTime = false;
		if (timeMinsInput === '') {
			submitTime = timer.state.timeReadSec;
		} else {
			submitTime = parseInt(timeMinsInput) * 60;
		}
	}
	// Reading pace calculations
	let pagesRead = $derived(estimateEndPage - actualStartPage);
	let pagesPerHour = $derived(submitTime !== 0 ? Math.floor(pagesRead / (submitTime / 3600)) : 0);
	let readingPerformance = $derived.by(() => {
		return getPaceSpeed(pagesPerHour);
	});
	// Start / End page nr in case of edit and validation
	let placeholderStart = actualStartPage.toString();
	let startInput = $state(placeholderStart);
	let placeholderEnd = $derived(estimateEndPage.toString());
	let endInput = $derived(placeholderEnd);
	function validateInput(isStart: boolean) {
		if (startInput === '') startInput = '0';
		if (endInput === '') endInput = '0';

		const startPage = parseInt(startInput);
		const endPage = parseInt(endInput);
		if (startPage > endPage) {
			if (isStart) {
				endInput = startInput;
			} else {
				startInput = endInput;
			}
		}
		pagesRead = endPage - startPage;
		if (pagesRead < 0) pagesRead = 0;
	}
	// Set goal time and modal toggle
	const goalOptions = [0.5, 30, 45, 60, 75, 90, 120, 150, 180]; //minutes
	let openGoalModal = $state(false);
	async function setGoal(goalMin: number) {
		openGoalModal = false;
		ts.goalTime = goalMin * 60;
		await new Promise((resolve) => setTimeout(resolve, 200));
		ts.isGoalSet = true;
	}
	// Change time from seconds to H:MM:SS format
	let formatedTime: string = $derived.by(() => formatTime(ts.timeReadSec));
	function formatTime(durationSec: number): string {
		const hour = Math.floor(durationSec / 3600);
		const min = Math.floor(durationSec / 60 - hour * 60);
		const sec = Math.floor(durationSec - hour * 3600 - min * 60);
		return `${hour.toString().padStart(1, '0')}:${min.toString().padStart(2, '0')}:${sec.toString().padStart(2, '0')}`;
	}
	// Done reading opens end session modal with stats
	let openSessionModal = $state(false);
	function endSession() {
		timer.state.isRunning = false; // pause timer
		openSessionModal = true;
		timeMinsInput = '';
		submitTime = timer.state.timeReadSec;
		const pace = isFinite(actualPace) ? actualPace : 0;
		estimateEndPage = actualStartPage + Math.floor((submitTime / 60) * pace);
	}
</script>

{#snippet selectTimeBtn(amountMin: number)}
	<button
		class="py-1 hover:bg-current/15 focus:bg-current/15 focus:outline-0"
		onclick={() => setGoal(amountMin)}>{formatTime(amountMin * 60)}</button
	>
{/snippet}

<div class="relative flex h-fit w-[332px] flex-col p-4 sm:w-[532px]">
	<div class="absolute size-[300px] rounded-full bg-current/15 p-0.5 sm:size-[500px]">
		<div
			class="text-light dark:text-dark size-full rounded-full"
			style="background: conic-gradient(
            {timer.overreadLoop % 2 ? 'transparent' : 'currentColor'} 0% {timer.loopProgress}%, 
            {timer.overreadLoop % 2 ? 'currentColor' : 'transparent'} {timer.loopProgress}%)"
		></div>
	</div>
	<div
		class="border-light dark:border-dark z-10 flex size-[300px] justify-center rounded-full transition-colors sm:size-[500px]
        {!ts.isGoalSet
			? 'bg-logo-blue/20'
			: !ts.isTimerStarted
				? 'bg-logo-purple/20'
				: ts.isRunning
					? timer.complete
						? 'bg-status-logo-done/20'
						: 'bg-status-logo-progress/20'
					: 'bg-logo-purple/20'}"
	>
		<div class="relative flex flex-col justify-center text-center">
			<div class="w-full font-sans text-2xl text-current/80 tabular-nums sm:text-4xl">
				{formatedTime}
			</div>
			{#if ts.isGoalSet}
				{#if !ts.isTimerStarted}
					<button
						class="text-dark active:inset-shadow-md bg-logo-purple mx-auto w-full rounded-2xl p-1 hover:brightness-95 focus:outline sm:p-2 sm:text-lg"
						onclick={() => timer.startTimer()}>Start</button
					>
				{:else}
					<div
						class="w-full font-sans text-2xl text-current/20 tabular-nums sm:text-4xl"
						in:slide={{ axis: 'x' }}
					>
						{formatTime(ts.goalTime)}
					</div>
				{/if}
			{:else}
				<div class="relative mx-auto w-full">
					<button
						class="text-dark active:inset-shadow-md bg-logo-blue w-full rounded-2xl p-1 hover:brightness-95 focus:outline sm:p-2 sm:text-lg"
						onclick={() => (openGoalModal = true)}>Set goal</button
					>
					{#if openGoalModal}
						<div
							class="fixed inset-0 z-10"
							aria-hidden="true"
							onclick={() => (openGoalModal = false)}
						></div>
						<div
							role="dialog"
							tabindex="-1"
							aria-modal="true"
							class="absolute z-10 mt-1 max-h-[99px] w-full overflow-y-scroll rounded-lg sm:max-h-[132px]"
							transition:slide={{ duration: 200 }}
						>
							<div class="flex flex-col divide-y divide-current/15 bg-current/15">
								{#each goalOptions as goalTime}
									{@render selectTimeBtn(goalTime)}
								{/each}
							</div>
						</div>
					{/if}
				</div>
			{/if}
		</div>
	</div>
	{#if ts.isTimerStarted}
		<div class="text-dark mt-6 flex w-full justify-around" transition:slide={{ delay: 300 }}>
			<button
				class="active:inset-shadow-md mx-4 w-full rounded-2xl p-2 text-lg
            transition-colors hover:brightness-95 active:text-base {ts.isRunning
					? 'bg-status-logo-progress'
					: 'bg-logo-purple'}"
				onclick={() => timer.togglePauseTimer()}
				>{#if ts.isRunning}Pause{:else}Resume{/if}</button
			>
			<button
				class="active:inset-shadow-md bg-status-logo-done mx-4 w-full rounded-2xl p-2 text-lg hover:brightness-95 active:text-base"
				onclick={endSession}>Done</button
			>
		</div>
	{/if}

	{#if openSessionModal}
		<div
			class="bg-light dark:bg-dark absolute z-10 h-[382px] w-[300px] overflow-hidden rounded-xl sm:h-[584px] sm:w-[500px]"
			transition:slide
		>
			<form
				method="POST"
				use:enhance
				class="flex size-full flex-col gap-2 bg-current/15 px-2 pb-3.5 sm:gap-4 sm:p-4"
			>
				<PageSubheader>Reading session</PageSubheader>
				<div class="flex justify-between gap-2 text-current/80 sm:gap-4">
					<div class="indented bg-logo-purple/20 flex grow flex-col justify-around">
						<div class="border-b border-dashed pb-1">
							Start page:
							<span class="float-right">
								<input type="hidden" name="firstPage" value={progress.first_page} />
								<input
									class="outline-0 focus:underline"
									name="start"
									type="text"
									inputmode="numeric"
									placeholder={placeholderStart}
									bind:value={startInput}
									oninput={() => (startInput = startInput.replace(/\D/g, ''))}
									onblur={() => validateInput(true)}
									maxlength="4"
									required
								/>
								<span class="icon-[solar--pen-outline]"></span>
							</span>
						</div>
						<div>
							<label for="end">End page:</label>
							<span class="float-right">
								<input
									class="outline-0 focus:underline"
									name="end"
									type="text"
									inputmode="numeric"
									placeholder={placeholderEnd}
									bind:value={endInput}
									oninput={() => (endInput = endInput.replace(/\D/g, ''))}
									onblur={() => validateInput(false)}
									maxlength="4"
									required
								/>
								<input type="hidden" name="pagesRead" value={progress.pages_read + pagesRead} />
								<span class="icon-[solar--pen-outline]"></span>
							</span>
						</div>
					</div>
					<div class="indented hidden grow sm:block"></div>
					<div class="indented bg-status-logo-progress/20">
						<h1 class="font-serif text-4xl text-shadow-lg sm:text-5xl">{pagesRead}</h1>
						<span class="text-sm italic sm:text-base">pages read</span>
					</div>
				</div>
				<div
					class="indented w-full text-center {editTime
						? 'bg-logo-purple/20'
						: submitTime >= timer.state.goalTime
							? 'bg-status-logo-done/20'
							: 'bg-logo-red/20'}"
				>
					<input name="time" type="hidden" value={progress.time_read + submitTime} />
					Time read:
					<span class="font-semibold">
						{#if !editTime}
							<div class="inline-flex h-5">
								{formatTime(submitTime)}
								<button onclick={() => (editTime = true)} aria-label="Edit time">
									<span class="icon-[solar--pen-2-bold] size-5"></span>
								</button>
							</div>
						{:else}
							<!-- svelte-ignore a11y_autofocus -->
							<input
								class="time outline-0 focus:underline"
								type="text"
								inputmode="numeric"
								placeholder={'~' + Math.round(submitTime / 60)}
								bind:value={timeMinsInput}
								onblur={() => checkTimeInput()}
								oninput={() => (timeMinsInput = timeMinsInput.replace(/\D/g, ''))}
								maxlength="3"
								autofocus
							/>min.
						{/if}
						/ {formatTime(timer.state.goalTime)}</span
					>
				</div>
				<div class="flex flex-row gap-2 sm:gap-4">
					<div class="indented w-3/4 text-center">
						Pages per hour: <span class="font-semibold">~{pagesPerHour}</span>
					</div>
					<div
						class="indented grow text-center font-semibold text-current/70 capitalize italic
						{readingPerformance === 'fast'
							? 'bg-status-logo-done/20'
							: readingPerformance === 'average'
								? 'bg-status-logo-progress/20'
								: 'bg-logo-red/20'}"
					>
						{readingPerformance}
					</div>
				</div>

				<div class="text-dark mt-auto flex justify-around">
					<button
						class="active:inset-shadow-md bg-logo-red mx-4 w-full rounded-2xl p-2 text-lg hover:brightness-95 active:text-base"
						type="button"
						onclick={() => (openSessionModal = false)}
					>
						Cancel
					</button>
					<button
						class="active:inset-shadow-md bg-status-logo-done mx-4 w-full rounded-2xl p-2 text-lg hover:brightness-95 active:text-base"
						type="submit"
					>
						Confirm
					</button>
				</div>
			</form>
		</div>
	{/if}
</div>

<style>
	input {
		width: 3rem;
		padding: 0px 2px;
		border-width: 0;
		border-radius: 0;
	}
	input.time {
		width: 2.3rem;
	}
</style>
