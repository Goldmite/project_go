import { PUBLIC_API_URL } from '$env/static/public';
import { error } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ params }) => {
	if (params.sharedId) {
		if (params.from === 'personal') {
			throw error(404, 'Personal cannot have an ID.');
		}
	} else {
		if (params.from === 'shared') {
			throw error(404, 'Shared requires an ID for specific group.');
		}
	}

	const res = await fetch(`${PUBLIC_API_URL}/books/${params.id}`);
	if (res.status != 200) {
		return error(res.status);
	}
	const book = await res.json();
	return { book };
};
