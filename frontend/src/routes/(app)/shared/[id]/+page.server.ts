import { get } from 'svelte/store';
import type { PageServerLoad } from './$types';
import { groups } from '$lib/stores/group';
import { error } from '@sveltejs/kit';
import { PUBLIC_API_URL } from '$env/static/public';

export const load: PageServerLoad = async ({ params }) => {
	const currGroup = get(groups).find((g) => g.id === params.id);
	if (!currGroup) {
		error(404, 'Not found');
	}
	const groupRes = await fetch(`${PUBLIC_API_URL}/books/groups/${currGroup.id}`);
	const books = await groupRes.json();

	const membersRes = await fetch(`${PUBLIC_API_URL}/users/groups/${currGroup.id}`);
	const members = await membersRes.json();
	return { books, currGroup, members };
};
