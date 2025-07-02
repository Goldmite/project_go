import { fail, redirect } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';
import { get } from 'svelte/store';
import { user } from '$lib/stores/user';
import { groups } from '$lib/stores/group';
import { PUBLIC_API_URL } from '$env/static/public';

export const load: LayoutServerLoad = async () => {
	const loggedInUser = get(user);
	if (!loggedInUser) {
		redirect(307, '/login');
	}

	const res = await fetch(`${PUBLIC_API_URL}/groups/${loggedInUser?.id}`);
	if (res.status != 200) {
		return fail(res.status);
	}

	const shared = await res.json();

	groups.set(shared);
	return { shared };
};
