import { redirect } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';
import { get } from 'svelte/store';
import { user } from '$lib/stores/user';

export const load: LayoutServerLoad = () => {
    const loggedInUser = get(user);
	if (!loggedInUser) {
		redirect(307, '/login');
	}
};