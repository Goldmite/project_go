import { get } from 'svelte/store';
import type { PageServerLoad } from '../../personal/$types';
import { user } from '$lib/stores/user';
import { PUBLIC_API_URL } from '$env/static/public';
import { fail, redirect, type Actions } from '@sveltejs/kit';

export const load: PageServerLoad = async () => {
	const userId = get(user)?.id;
	if (!userId) return;

	const invsRes = await fetch(`${PUBLIC_API_URL}/users/invites/${userId}`);
	if (invsRes.status != 200) {
		return fail(invsRes.status);
	}
	const invitations = await invsRes.json();
	return { invitations };
};

export const actions = {
	accept: async (event) => {
		const currUser = get(user);
		const formData = await event.request.formData();
		const groupId = formData.get('group_id');
		const toRedirect: boolean = formData.get('redirect') === 'true';
		if (!currUser || !groupId) return fail(400);
		if (toRedirect) {
			redirect(303, `/shared/redirect/${groupId}`);
		}

		const data = JSON.stringify({
			user_id: currUser.id,
			email: currUser.email,
			group_id: groupId
		});
		const acceptRes = await fetch(`${PUBLIC_API_URL}/groups/invites/accept`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: data
		});

		if (acceptRes.status != 200) {
			return fail(acceptRes.status);
		}
		return { success: true, isAccepted: true };
	},
	decline: async (event) => {
		const userEmail = get(user)?.email;
		const groupId = (await event.request.formData()).get('group_id');
		if (!userEmail || !groupId) return fail(400);

		const data = JSON.stringify({
			email: userEmail,
			group_id: groupId
		});
		const declineRes = await fetch(`${PUBLIC_API_URL}/groups/invites/decline`, {
			method: 'DELETE',
			headers: { 'Content-Type': 'application/json' },
			body: data
		});

		if (declineRes.status != 200) {
			return fail(declineRes.status);
		}
		return { success: true, isAccepted: false };
	}
} satisfies Actions;
