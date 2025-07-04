import { getContext, setContext } from 'svelte';
import type { User } from './types/user';

export class UserState {
	user = $state<User>();

	constructor(loggedIn: User) {
		this.user = loggedIn;
	}
}

const USER_KEY = Symbol('USER_LOGIN');

export function setUserState(loggedIn: User) {
	return setContext(USER_KEY, new UserState(loggedIn));
}

export function getUserState() {
	return getContext<ReturnType<typeof setUserState>>(USER_KEY);
}
