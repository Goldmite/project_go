import { user } from '$lib/stores/user';
import { get } from 'svelte/store';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async () => {
  return { user: get(user) ?? null};
};
