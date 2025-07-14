export function formatMinToHoursOrMin(min: number): string {
	const h = Math.floor(min / 60);
	const m = min > 0 ? Math.floor(min - h * 60) : 0;
	return h > 0 ? `${h}h${m < 10 ? '0' : ''}${m}m` : `${m} min`;
}
