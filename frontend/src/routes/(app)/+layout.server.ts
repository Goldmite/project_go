import { fail, redirect } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';
import { get } from 'svelte/store';
import { user } from '$lib/stores/user';
import { groups } from '$lib/stores/group';
import type { Group } from '$lib/types/group';

export const load: LayoutServerLoad = async () => {
	const loggedInUser = get(user);
	if (!loggedInUser) {
		redirect(307, '/login');
	}

	const res = await fetch(`http://localhost:3000/api/groups/${loggedInUser?.id}`);
	if (res.status != 200) {
		return fail(res.status);
	}

	const groupsData = await res.json();
	const shared = groupsData.map((g: Group, index: number) => ({
		...g,
		name: `${g.name}-${index + 1}`
	}));

	groups.set(shared);
	return { shared };
};
