import type { User } from '$lib/types/user';
import { writable } from 'svelte/store';

function createLocalStore<T>(key: string, initial: T) {
	let data = initial;
	if (typeof localStorage !== 'undefined') {
		const stored = localStorage.getItem(key);
		data = stored ? JSON.parse(stored) : initial;
	}
	const store = writable<T>(data);
	if (typeof localStorage !== 'undefined') {
		store.subscribe((value) => {
			localStorage.setItem(key, JSON.stringify(value));
		});
	}

	return store;
}

export const bookOwners = createLocalStore<string[]>('book', []);

export const groupMembers = createLocalStore<User[]>('members', []);
