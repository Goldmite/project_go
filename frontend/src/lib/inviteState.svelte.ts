import { getContext, setContext } from 'svelte';
import type { User } from './types/user';

export class InviteState {
	invites = $state<User[]>([]);

	constructor() {}

	add(id: string, name: string, email: string) {
		this.invites.push({
			id,
			name,
			email
		});
	}

	remove(id: string) {
		this.invites = this.invites.filter((i) => i.id !== id);
	}

	isExist(id: string) {
		return this.invites.some((i) => i.id === id);
	}
}

const INVITE_KEY = Symbol('INVITE');

export function setInviteState() {
	return setContext(INVITE_KEY, new InviteState());
}

export function getInviteState() {
	return getContext<ReturnType<typeof setInviteState>>(INVITE_KEY);
}
