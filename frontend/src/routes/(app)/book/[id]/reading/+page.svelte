<script lang="ts">
	import { slide } from 'svelte/transition';

	let openGoalModal = $state(false);
	let isGoalSet = $state(false);
	let isTimerStarted = $state(false);
	let isRunning = $state(false);
	function startTimer() {
		isTimerStarted = true;
		isRunning = true;
	}
	function togglePauseTimer() {
		isRunning = !isRunning;
	}

	const goalOptions = [30, 45, 60, 75, 90, 120, 150, 180]; //minutes
	let goalTime = $state(30 * 60);
	async function setGoal(goalMin: number) {
		openGoalModal = false;
		goalTime = goalMin * 60;
		await new Promise((resolve) => setTimeout(resolve, 200));
		isGoalSet = true;
	}
	let timeReadSec: number = $state(0);
	let complete = $derived(timeReadSec >= goalTime);
	let readProgress = $derived((timeReadSec / goalTime) * 100); // [0 - x], where x is 100 * loop number
	let overreadLoop = $derived(Math.floor(readProgress / 100)); // [0 - y], where y is the amount of times the loop progress reached 100
	let loopProgress = $derived(readProgress - overreadLoop * 100); // [0 - 100]

	let formatedTime: string = $derived.by(() => formatTime(timeReadSec));
	let readingInterval: number | undefined;
	$effect(() => {
		if (isRunning) {
			readingInterval = setInterval(() => {
				timeReadSec = timeReadSec + 1;
			}, 1000);
		} else {
			clearInterval(readingInterval);
		}
	});
	function formatTime(durationSec: number): string {
		const hour = Math.floor(durationSec / 3600);
		const min = Math.floor(durationSec / 60 - hour * 60);
		const sec = Math.floor(durationSec - hour * 3600 - min * 60);
		return `${hour.toString().padStart(1, '0')}:${min.toString().padStart(2, '0')}:${sec.toString().padStart(2, '0')}`;
	}
</script>

{#snippet selectTimeBtn(amountMin: number)}
	<button
		class="py-1 hover:bg-current/15 focus:bg-current/15 focus:outline-0"
		onclick={() => setGoal(amountMin)}>{formatTime(amountMin * 60)}</button
	>
{/snippet}

<div class="relative flex h-fit w-[332px] flex-col rounded-xl p-4 sm:w-[532px]">
	<div class="absolute size-[300px] rounded-full bg-current/15 p-0.5 sm:size-[500px]">
		<div
			class="text-light dark:text-dark size-full rounded-full"
			style="background: conic-gradient(
            {overreadLoop % 2 ? 'transparent' : 'currentColor'} 0% {loopProgress}%, 
            {overreadLoop % 2 ? 'currentColor' : 'transparent'} {loopProgress}%)"
		></div>
	</div>
	<div
		class="border-light dark:border-dark z-10 flex size-[300px] justify-center rounded-full transition-colors sm:size-[500px]
        {!isGoalSet
			? 'bg-logo-blue/20'
			: !isTimerStarted
				? 'bg-logo-purple/20'
				: isRunning
					? complete
						? 'bg-status-logo-done/20'
						: 'bg-status-logo-progress/20'
					: 'bg-logo-purple/20'}"
	>
		<div class="relative flex flex-col justify-center text-center">
			<div class="w-full font-sans text-2xl text-current/80 tabular-nums sm:text-4xl">
				{formatedTime}
			</div>
			{#if isGoalSet}
				{#if !isTimerStarted}
					<button
						class="text-dark active:inset-shadow-md bg-logo-purple mx-auto w-full rounded-2xl p-1 hover:brightness-95 focus:outline sm:p-2 sm:text-lg"
						onclick={() => startTimer()}>Start</button
					>
				{:else}
					<div
						class="w-full font-sans text-2xl text-current/20 tabular-nums sm:text-4xl"
						in:slide={{ axis: 'x' }}
					>
						{formatTime(goalTime)}
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
	{#if isTimerStarted}
		<div class="text-dark mt-6 flex w-full justify-around" transition:slide={{ delay: 300 }}>
			<button
				class="active:inset-shadow-md mx-4 w-full rounded-2xl p-2 text-lg
            transition-colors hover:brightness-95 active:text-base {isRunning
					? 'bg-status-logo-progress'
					: 'bg-logo-purple'}"
				onclick={togglePauseTimer}
				>{#if isRunning}Pause{:else}Resume{/if}</button
			>
			<button
				class="active:inset-shadow-md bg-status-logo-done mx-4 w-full rounded-2xl p-2 text-lg hover:brightness-95 active:text-base"
				>Done</button
			>
		</div>
	{/if}
</div>
