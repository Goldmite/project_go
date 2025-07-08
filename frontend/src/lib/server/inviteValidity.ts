import { PUBLIC_API_URL } from '$env/static/public';
import type { User } from '$lib/types/user';
import { fail } from '@sveltejs/kit';

export async function checkUser(email: FormDataEntryValue, currUsr: User | null) {
	if (email === '') return fail(400);
	const currUserEmail = currUsr?.email ?? '';
	if (currUserEmail == email) {
		return fail(409, { email, yourself: true });
	}

	const res = await fetch(`${PUBLIC_API_URL}/users/${email}`);

	if (res.status == 404) {
		return fail(404, { email, notfound: true });
	}
	if (res.status != 200) {
		return fail(res.status);
	}
	const invitee = await res.json();
	return fail(400, { invitee, fakesuccess: true });
}
