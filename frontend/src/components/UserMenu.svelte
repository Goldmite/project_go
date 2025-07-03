<script lang="ts">
	import { enhance } from "$app/forms";
	import type { User } from "$lib/types/user";

    let { userInfo }: {userInfo: User } = $props();
</script>

<div class="absolute z-10 h-fit m-2">
    {userInfo.name}
    {userInfo.email}
    <form method='POST' action='/logout' use:enhance>
    <button class="border"
        type="submit">Log out</button>
</form>
</div>