import { fail, type Actions } from '@sveltejs/kit';
import VITE_BASE_API_URL from '$env/static/public'

export const actions = {
	login: async (event) => {
		// TODO log the user in
	},
	signup: async (event) => {
		// TODO register the user
		const data = await event.request.formData();
		
		const res = await fetch('http://localhost:3000/api/users', {
			method: 'POST',
			body: data,
		});
		if (res.status != 201) {
			return fail(res.status, { res });
		}
		console.log(res);
		return { success: true };
	}
} satisfies Actions;