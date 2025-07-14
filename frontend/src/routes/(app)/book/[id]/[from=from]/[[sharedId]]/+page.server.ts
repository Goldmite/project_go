import { PUBLIC_API_URL } from '$env/static/public';
import { error } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ params, locals }) => {
	if (params.sharedId) {
		if (params.from === 'personal') {
			throw error(404, 'Personal cannot have an ID.');
		}
	} else {
		if (params.from === 'shared') {
			throw error(404, 'Shared requires an ID for specific group.');
		}
	}

	const detailsRes = await fetch(`${PUBLIC_API_URL}/books/${params.id}`);
	if (detailsRes.status != 200) {
		return error(detailsRes.status);
	}
	const userId = locals.user?.id;
	let progress = undefined;
	if (userId) {
		const req = new URLSearchParams({ user_id: userId, isbn: params.id });
		const progressRes = await fetch(`${PUBLIC_API_URL}/stats/progress/book?${req.toString()}`);
		if (progressRes.status == 200) {
			progress = await progressRes.json();
		}
	}
	const book = await detailsRes.json();
	return { book, progress };
};
