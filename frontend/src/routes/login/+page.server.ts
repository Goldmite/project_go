import { fail, redirect, type Actions } from '@sveltejs/kit';
import { user } from '$lib/stores/user';
import { PUBLIC_API_URL } from '$env/static/public';

export const actions = {
	login: async (event) => {
		const data = await event.request.formData();

		const res = await fetch(`${PUBLIC_API_URL}/users/login`, {
			method: 'POST',
			body: data
		});

		if (res.status != 200) {
			return fail(res.status);
		}

		const userJson = await res.json();
		user.set(userJson);
		redirect(303, '/personal');
	},
	signup: async (event) => {
		const data = await event.request.formData();

		const res = await fetch(`${PUBLIC_API_URL}/users/signup`, {
			method: 'POST',
			body: data
		});

		if (res.status != 201) {
			return fail(res.status);
		}

		return { success: true };
	}
} satisfies Actions;
