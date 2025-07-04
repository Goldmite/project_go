import { user } from '$lib/stores/user';
import { redirect, type Actions } from '@sveltejs/kit';

export const actions = {
	default: async () => {
		user.set(null);
		redirect(303, '/login');
	}
} satisfies Actions;
