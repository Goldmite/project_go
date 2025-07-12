import { PUBLIC_API_URL } from '$env/static/public';
import { fail, redirect } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ params, locals }) => {
	const isbn = params.id;
	const userId = locals.user?.id;
	if (!userId) return;

	const req = new URLSearchParams({ user_id: userId, isbn: isbn });
	const res = await fetch(`${PUBLIC_API_URL}/stats/progress/book?${req.toString()}`);

	if (res.status != 200) {
		return fail(res.status);
	}
	const progress = await res.json();
	return { progress };
};

export const actions = {
	default: async (event) => {
		const isbn = event.params.id;
		const userId = event.locals.user?.id;
		if (!userId) return;
		const data = await event.request.formData();
		const firstPage = data.get('firstPage');
		const startPage = data.get('start');
		const pagesRead = data.get('pagesRead');
		const timeRead = data.get('time');
		if (firstPage == '' || startPage == '' || pagesRead == '' || timeRead == '') {
			return fail(400);
		}
		const firstRealPage = Number(firstPage) === 1 ? Number(startPage) : undefined;
		const req = JSON.stringify({
			user_id: userId,
			isbn: isbn,
			pages_read: Number(pagesRead),
			time_read: Number(timeRead),
			first_page: firstRealPage
		});
		const res = await fetch(`${PUBLIC_API_URL}/stats/progress/book`, {
			method: 'PUT',
			headers: { 'Content-Type': 'application/json' },
			body: req
		});
		if (res.status != 200) {
			return fail(res.status);
		}

		redirect(303, `/book/${isbn}`);
	}
} satisfies Actions;
