import { fail, type Actions } from '@sveltejs/kit';
import VITE_BASE_API_URL from '$env/static/public'

export const actions = {
	login: async (event) => {
		const data = await event.request.formData();
		
		const res = await fetch('http://localhost:3000/api/users/login', {
			method: 'POST',
			body: data
		});

		if (res.status != 200) {
			return fail(res.status, { res });
		}

		return { success: true };
	},
	signup: async (event) => {
		const data = await event.request.formData();
		
		const res = await fetch('http://localhost:3000/api/users/signup', {
			method: 'POST',
			body: data,
		});

		if (res.status != 201) {
			return fail(res.status, { res });
		}

		return { success: true };
	}
} satisfies Actions;