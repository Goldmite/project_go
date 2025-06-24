import { get } from "svelte/store";
import type { PageServerLoad } from "../personal/$types";
import { user } from '$lib/stores/user';
import { groups } from '$lib/stores/group';
import type { Group } from '$lib/types/group';
import { fail } from "@sveltejs/kit";

export const load: PageServerLoad = async () => {
    const currUser = get(user);

    const res = await fetch(`http://localhost:3000/api/groups/${currUser?.id}`)
    if (res.status != 200) {
        return fail(res.status);
    }

    const json = await res.json();
    const groupsData = json.map((g: Group, index: number) => ({
        ...g,
        name: `${g.name}-${index + 1}`
    }));

    groups.set(groupsData);
    return { groupsData }
};