export function toRoman(nr: number): string {
	if (nr >= 100) return '';

	const numeralsMap = [
		['I', 'II', 'III', 'IV', 'V', 'VI', 'VII', 'VIII', 'IX'],
		['X', 'XX', 'XXX', 'XL', 'L', 'LX', 'LXX', 'LXXX', 'XC']
	];

	const digits = nr.toString().split('');
	let position = digits.length - 1;

	return digits.reduce((roman, digit) => {
		if (digit !== '0') {
			roman += numeralsMap[position][parseInt(digit) - 1];
		}
		position -= 1;
		return roman;
	}, '');
}
