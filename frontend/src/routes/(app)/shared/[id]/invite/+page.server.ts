import { checkUser } from '$lib/server/inviteValidity';
import { error, fail, redirect, type Actions } from '@sveltejs/kit';
import { get } from 'svelte/store';
import type { PageServerLoad } from './$types';
import { groups } from '$lib/stores/group';
import { PUBLIC_API_URL } from '$env/static/public';

export const load: PageServerLoad = async ({ params }) => {
	const currGroup = get(groups).find((g) => g.id === params.id);
	if (!currGroup) {
		error(404, 'Not found');
	}
	const sentRes = await fetch(`${PUBLIC_API_URL}/groups/invites/${currGroup.id}`);
	if (sentRes.status != 200) {
		return fail(sentRes.status);
	}
	const invitedUsers = await sentRes.json();
	if (invitedUsers === null) return undefined;
	return { invitedUsers };
};
export const actions = {
	invite: async (event) => {
		const emails = (await event.request.formData()).getAll('emails[]');
		if (!emails) return fail(404);
		const userId = event.locals.user?.id ?? '';
		const groupId = event.params.id ?? '';

		const inviteForm = new FormData();
		emails.forEach((email) => {
			inviteForm.append('email_to', email);
		});
		inviteForm.append('group_id', groupId);
		inviteForm.append('invited_by', userId);
		const inviteRes = await fetch(`${PUBLIC_API_URL}/groups/invites`, {
			method: 'POST',
			body: inviteForm
		});
		if (inviteRes.status != 200) {
			return fail(inviteRes.status);
		}
		redirect(303, `/shared/${groupId}`);
	},
	checkUser: async (event) => {
		const email = (await event.request.formData()).get('email') ?? '';
		return checkUser(email, event.locals.user);
	}
} satisfies Actions;
