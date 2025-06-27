import { groups } from '$lib/stores/group';
import { user } from '$lib/stores/user';
import { fail, redirect, type Actions } from '@sveltejs/kit';
import { get } from 'svelte/store';


export const actions = {
	createShared: async (event) => {
		const data = (await event.request.formData());
		let name = data.get('name')?.toString() ?? '';
		if (name === '') name = 'Group'
		const userId = get(user)?.id ?? '';

		const createForm = new FormData();
		createForm.append('id', userId);
		createForm.append('name', name);

		const createRes = await fetch('http://localhost:3000/api/groups', {
			method: 'POST',
			body: createForm
		});

		if (createRes.status != 201) {
			return fail(createRes.status);
		}
		const groupId = await createRes.json();
		const emails = data.getAll('emails[]');
		const inviteForm = new FormData();
		emails.forEach((email) => {
			inviteForm.append('email_to', email);
		});
		inviteForm.append('group_id', groupId);
		inviteForm.append('invited_by', userId);
		const inviteRes = await fetch('http://localhost:3000/api/groups/invites', {
			method: 'POST',
			body: inviteForm
		});
		if (inviteRes.status != 200) {
			return fail(inviteRes.status);
		}
		// DOESNT WORK BECAUSE LAYOUT HASNT CREATED GROUP NAME URL YET
		const gotoGroup = get(groups).find((g) => g.id === groupId);
		if (gotoGroup) {
			redirect(303, `/shared/${gotoGroup?.name}`);
		}
		return { success: true, groupId };
	},
	checkUser: async (event) => {
		const email = (await event.request.formData()).get('email') ?? '';
		if (email === '') return fail(400);
		const currUserEmail = get(user)?.email;
		if (currUserEmail == email) {
			return fail(409, { email, duplicate: true });
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
	
} satisfies Actions;
