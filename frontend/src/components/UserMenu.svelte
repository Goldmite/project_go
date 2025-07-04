<script lang="ts">
	import { enhance } from "$app/forms";
	import type { User } from "$lib/types/user";
	import { fly, slide } from "svelte/transition";

    let { userInfo }: {userInfo: User } = $props();
    const [ local, domain ] = userInfo.email.split('@', 2);
</script>

<div role="dialog" tabindex="-1" aria-modal="true" 
    class="absolute right-1.5 sm:right-6 rounded-lg overflow-hidden 
    z-10 w-42 sm:w-56 h-fit bg-light dark:bg-dark"
    transition:slide={{ duration: 200 }}>
    <div class="flex flex-col gap-1 wrap-break-word w-full bg-current/15 p-4">
        <span class="font-semibold">{userInfo.name}</span>
        <div class="text-sm/4 italic text-wrap flex flex-wrap mb-2">
            <span>{local}</span><span>@{domain}</span>
        </div>
        <div class="flex gap-2 pt-2 border-t-1 border-dashed flex-col sm:flex-row w-full">
            <button class=" h-8 w-full sm:w-1/2 rounded-2xl bg-logo-blue text-dark
                active:inset-shadow-md  active:text-sm hover:brightness-90"
                type="button"
            >Edit</button>
            <form class="w-full sm:w-1/2" method='POST' action='/logout' use:enhance>
                <button class=" h-8 w-full rounded-2xl bg-logo-red text-dark
                    active:inset-shadow-md active:text-sm hover:brightness-90"
                    type="submit"
                >Log out</button>
            </form>
        </div>
    </div>
</div>