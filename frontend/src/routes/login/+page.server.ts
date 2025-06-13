import { fail, type Actions } from '@sveltejs/kit';
import { user } from '$lib/stores/user';

export const actions = {
	login: async (event) => {
		const data = await event.request.formData();
		
		const res = await fetch('http://localhost:3000/api/users/login', {
			method: 'POST',
			body: data
		});

		if (res.status != 200) {
			return fail(res.status);
		}

		const userJson = await res.json();
		user.set(userJson);

		return { success: true };
	},
	signup: async (event) => {
		const data = await event.request.formData();
		
		const res = await fetch('http://localhost:3000/api/users/signup', {
			method: 'POST',
			body: data,
		});

		if (res.status != 201) {
			return fail(res.status);
		}

		return { success: true };
	}
} satisfies Actions;