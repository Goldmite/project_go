export function getPaceSpeed(pph: number): string {
	if (pph < 50) {
		return 'slow';
	} else if (pph <= 70) {
		return 'average';
	} else {
		return 'fast';
	}
}
