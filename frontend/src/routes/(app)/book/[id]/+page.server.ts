import { PUBLIC_API_URL } from '$env/static/public';
import { error } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ params }) => {
	const res = await fetch(`${PUBLIC_API_URL}/books/${params.id}`);
	if (res.status != 200) {
		return error(res.status);
	}
	const book = await res.json();
	return { book };
};
