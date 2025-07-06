enum Language {
	en = 'English',
	lt = 'Lithuanian'
}

export function toReadableLanguage(code: string): string {
	const langName = Language[code as keyof typeof Language] || '_';
	return langName;
}
