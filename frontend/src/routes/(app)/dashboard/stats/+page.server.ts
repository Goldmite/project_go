import { PUBLIC_API_URL } from '$env/static/public';
import { fail } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import type { Stats } from '$lib/types/responses/stats';

export const load: PageServerLoad = async ({ locals }) => {
	const userId = locals.user?.id;
	if (!userId) return;

	const statsRes = await fetch(`${PUBLIC_API_URL}/stats/user/${userId}`);
	if (statsRes.status != 200) {
		if (statsRes.status == 404) {
			const total: Stats = { total_pages: 0, total_time: 0 };
			return { total };
		}
		return fail(statsRes.status);
	}
	const total = await statsRes.json();
	return { total };
};
