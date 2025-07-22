import { PUBLIC_API_URL } from '$env/static/public';
import { error } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async ({ locals }) => {
	const userId = locals.user?.id;
	if (!userId) return;

	const defaultFrom = '2025-01-01';
	const req = new URLSearchParams({ user_id: userId, from: defaultFrom });

	const activityRes = await fetch(`${PUBLIC_API_URL}/stats/activity?${req.toString()}`);
	const recentRes = await fetch(`${PUBLIC_API_URL}/books/recent/user/${userId}`);

	if (activityRes.status != 200) {
		return error(activityRes.status);
	}
	if (recentRes.status != 200) {
		return error(recentRes.status);
	}

	const activity = await activityRes.json();
	const recent = await recentRes.json();
	return { activity, recent };
};
