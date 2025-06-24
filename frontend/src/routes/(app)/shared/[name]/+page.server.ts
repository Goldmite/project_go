import { get } from "svelte/store";
import type { PageServerLoad } from "./$types";
import { groups } from "$lib/stores/group";

export const load: PageServerLoad = async ({ params }) => {
	const currGroup = get(groups).find(g => g.name === params.name);

	const groupRes = await fetch(`http://localhost:3000/api/books/groups/${currGroup?.id}`);
	const books = await groupRes.json();

	const membersRes = await fetch(`http://localhost:3000/api/users/groups/${currGroup?.id}`);
	const members = await membersRes.json();
	return { books, currGroup, members };
};