import type { PageServerLoad } from './$types';
import { user } from '$lib/stores/user';
import { get } from 'svelte/store';

export const load: PageServerLoad = async () => {
	const currUser = get(user);
	if (!currUser) {
		console.log("No user, redirect to login page");
	}
	const res = await fetch(`http://localhost:3000/api/books/user/${currUser?.id}`);
	const json = await res.json();
	
	return { json };
};