import { user } from '$lib/stores/user';
import { fail } from '@sveltejs/kit';
import { get } from 'svelte/store';

export async function checkUser(email: FormDataEntryValue) {
	if (email === '') return fail(400);
	const currUserEmail = get(user)?.email;
	if (currUserEmail == email) {
		return fail(409, { email, yourself: true });
	}

	const res = await fetch(`http://localhost:3000/api/users/${email}`);

	if (res.status == 404) {
		return fail(404, { email, notfound: true });
	}
	if (res.status != 200) {
		return fail(res.status);
	}
	const invitee = await res.json();
	return fail(400, { invitee, fakesuccess: true });
}
