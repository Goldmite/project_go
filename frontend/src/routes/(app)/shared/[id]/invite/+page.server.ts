import { checkUser } from '$lib/server/checkUser';
import { user } from '$lib/stores/user';
import { fail, type Actions } from '@sveltejs/kit';
import { get } from 'svelte/store';

export const actions = {
    invite: async (event) => {
        const emails = (await event.request.formData()).getAll('emails[]');
        if (!emails) return fail(404);
		const userId = get(user)?.id ?? '';
        const groupId = event.params.id ?? '';

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

        return { success: true };
    },
    checkUser: async (event) => {
        const email = (await event.request.formData()).get('email') ?? '';
        return checkUser(email);
    }
} satisfies Actions;