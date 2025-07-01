// See https://svelte.dev/docs/kit/types#app.d.ts
// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		//interface Locals {}
		// interface PageData {}
		interface PageState {
			isLoginSelected: boolean;
			invitations;
			selected;
			invite;
			groupMembers: User[];
		}
		// interface Platform {}
	}
}

export {};
