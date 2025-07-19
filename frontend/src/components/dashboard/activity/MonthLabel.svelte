<script lang="ts">
	let { gridCol, usedMonths }: { gridCol: any; usedMonths: string[] } = $props();
	const monthName = gridCol[0].date.split(' ')[1];
	let showName = $state('');
	if (usedMonths.length === 0 || usedMonths.at(-1) !== monthName) {
		if (gridCol[0].date.split(' ')[0] < 22) {
			usedMonths.push(monthName);
			showName = monthName;
		}
	}
</script>

<span class="text-xs/tight font-semibold text-current/80">{showName}</span>
