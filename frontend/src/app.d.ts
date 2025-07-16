// See https://svelte.dev/docs/kit/types#app.d.ts

import type { User } from '$lib/types/user';

// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		interface Locals {
			user: User | null;
		}
		// interface PageData {}
		interface PageState {
			isLoginSelected: boolean;
			invitations;
			stats;
			selected;
			invite;
			groupMembers: User[];
		}
		// interface Platform {}
	}
}

export {};
