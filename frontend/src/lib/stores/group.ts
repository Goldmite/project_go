import type { Group } from '$lib/types/group';
import { writable } from 'svelte/store';

export const groups = writable<Group[]>([]);
