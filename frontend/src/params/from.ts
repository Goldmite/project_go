import type { ParamMatcher } from '@sveltejs/kit';

export const match = ((param: string): param is ('personal' | 'shared') => {
	return param === 'personal' || param === 'shared';
}) satisfies ParamMatcher;