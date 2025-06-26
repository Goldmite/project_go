import { fail, type Actions } from "@sveltejs/kit";


export const actions = {
    checkUser: async (event) => {
        const email = (await event.request.formData()).get('email');

        const res = await fetch(`http://localhost:3000/api/users/${email}`);
        
        if (res.status == 404) {
			return fail(404, { email, notfound: true });
		}
        if (res.status != 200) {
			return fail(res.status);
		}
        const invitee = await res.json();
        return { success: true, invitee };
    }

} satisfies Actions;