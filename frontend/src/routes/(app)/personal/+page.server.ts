import type { PageServerLoad } from './$types';
import { user } from '$lib/stores/user';
import { get } from 'svelte/store';
import { fail, type Actions } from '@sveltejs/kit';
import { PUBLIC_API_URL } from '$env/static/public';

export const load: PageServerLoad = async () => {
	const currUser = get(user);
	const res = await fetch(`${PUBLIC_API_URL}/books/user/${currUser?.id}`);
	const books = await res.json();
	return { books };
};

export const actions = {
	default: async (event) => {
		const userId = get(user)?.id;
		const isbn = (await event.request.formData()).get('isbn');
		const len = isbn?.valueOf().toString().length;

		if (len != 10 && len != 13) {
			return fail(400, { isbn, incorrect: true });
		}

		const data = JSON.stringify({
			userId: userId,
			isbn: isbn
		});
		const res = await fetch(`${PUBLIC_API_URL}/books`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: data
		});

		if (res.status == 409) {
			return fail(409, { isbn, duplicate: true });
		}
		if (res.status == 404) {
			return fail(404, { isbn, notfound: true });
		}
		if (res.status != 200) {
			return fail(res.status);
		}

		return { success: true };
	}
} satisfies Actions;
