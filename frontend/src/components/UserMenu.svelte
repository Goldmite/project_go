<script lang="ts">
	import { enhance } from '$app/forms';
	import type { User } from '$lib/types/user';
	import { slide } from 'svelte/transition';

	let { userInfo }: { userInfo: User } = $props();
	const [local, domain] = userInfo.email.split('@', 2);
</script>

<div
	role="dialog"
	tabindex="-1"
	aria-modal="true"
	class="bg-light dark:bg-dark absolute right-1.5 z-10
    h-fit w-42 overflow-hidden rounded-lg sm:right-6 sm:w-56"
	transition:slide={{ duration: 200 }}
>
	<div class="flex w-full flex-col gap-1 bg-current/15 p-4 wrap-break-word">
		<span class="font-semibold">{userInfo.name}</span>
		<div class="mb-2 flex flex-wrap text-sm/4 text-wrap italic">
			<span>{local}</span><span>@{domain}</span>
		</div>
		<div class="flex w-full flex-col gap-2 border-t-1 border-dashed pt-2 italic sm:flex-row">
			<button
				class=" bg-logo-purple text-dark active:inset-shadow-md h-8 w-full rounded-2xl
                hover:brightness-90 active:text-sm sm:w-1/2"
				type="button">Edit</button
			>
			<form class="w-full sm:w-1/2" method="POST" action="/logout" use:enhance>
				<button
					class=" bg-logo-red text-dark active:inset-shadow-md h-8 w-full
                    rounded-2xl hover:brightness-90 active:text-sm"
					type="submit">Log out</button
				>
			</form>
		</div>
	</div>
</div>
