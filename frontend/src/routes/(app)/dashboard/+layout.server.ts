import { PUBLIC_API_URL } from '$env/static/public';
import { fail } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async ({ locals }) => {
	const userId = locals.user?.id;
	if (!userId) return;

	const defaultFrom = '2025-01-01';
	const req = new URLSearchParams({ user_id: userId, from: defaultFrom });

	const res = await fetch(`${PUBLIC_API_URL}/stats/activity?${req.toString()}`);
	if (res.status != 200) {
		return fail(res.status);
	}
	const activity = await res.json();
	return { activity };
};
