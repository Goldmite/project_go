export class Timer {
	state = $state({
		goalTime: 0,
		isGoalSet: false,
		isTimerStarted: false,
		isRunning: false,
		timeReadSec: 0
	});

	complete = $derived(this.state.timeReadSec >= this.state.goalTime);
	readProgress = $derived((this.state.timeReadSec / this.state.goalTime) * 100); // [0 - x], where x is 100 * loop number
	overreadLoop = $derived(Math.floor(this.readProgress / 100)); // [0 - y], where y is the amount of times the loop progress reached 100
	loopProgress = $derived(this.readProgress - this.overreadLoop * 100); // [0 - 100]

	#readingInterval: number | undefined;
	constructor() {
		$effect(() => {
			if (this.state.isRunning) {
				this.#readingInterval = setInterval(() => {
					this.state.timeReadSec = this.state.timeReadSec + 1;
				}, 1000);
			} else {
				clearInterval(this.#readingInterval);
			}
		});
	}

	startTimer() {
		this.state.isTimerStarted = true;
		this.state.isRunning = true;
	}
	togglePauseTimer() {
		this.state.isRunning = !this.state.isRunning;
	}
	resetTimer() {
		this.state = {
			goalTime: 0,
			isGoalSet: false,
			isTimerStarted: false,
			isRunning: false,
			timeReadSec: 0
		};
	}

	toJSON() {
		return {
			...this.state,
			isRunning: false
		};
	}
	static fromJSON(data: any) {
		return { ...data };
	}
}
