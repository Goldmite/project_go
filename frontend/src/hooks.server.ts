import { user } from '$lib/stores/user';
import type { Handle } from '@sveltejs/kit';
import { get } from 'svelte/store';

export const handle: Handle = async ({ event, resolve }) => {
	const currUsr = get(user);
	if (currUsr) {
		event.locals.user = currUsr;
	}
	const response = await resolve(event);

	return response;
};
