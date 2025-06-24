import { get } from "svelte/store";
import type { PageServerLoad } from "./$types";
import { groups } from "$lib/stores/group";

export const load: PageServerLoad = async ({ params }) => {
	const group = get(groups).find(g => g.name === params.name);

	const res = await fetch(`http://localhost:3000/api/books/groups/${group?.id}`);
	const books = await res.json();
	return { books, group };
};