import { get } from 'svelte/store';
import type { PageServerLoad } from './$types';
import { groups } from '$lib/stores/group';
import { error } from '@sveltejs/kit';

export const load: PageServerLoad = async ({ params }) => {
	const currGroup = get(groups).find((g) => g.id === params.id);
	if (!currGroup) {
		error(404, 'Not found');
	}
	const groupRes = await fetch(`http://localhost:3000/api/books/groups/${currGroup.id}`);
	const books = await groupRes.json();

	const membersRes = await fetch(`http://localhost:3000/api/users/groups/${currGroup.id}`);
	const members = await membersRes.json();
	return { books, currGroup, members };
};
