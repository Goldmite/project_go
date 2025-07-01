<script>
	import ActionButton from "./ActionButton.svelte";

    let { inv } = $props();
    let action = $state('');
</script>

<tr class="group border border-dark/40 font-sans tracking-tight ">
    <td class="italic px-1 border-r border-dark/10 group-hover:backdrop-brightness-95">{inv.group_name}</td>
    <td class="px-1 border-r border-dark/10 group-hover:backdrop-brightness-95">{inv.inviter_name}</td>
    <td class="p-1 group-hover:backdrop-brightness-95">
        <form method="POST" class="flex flex-row gap-2 sm:gap-4">
            <input type="hidden" name="group_id" value={inv.group_id} >
            <ActionButton toAccept={false} bind:action={action} />
            <ActionButton toAccept={true} bind:action={action} />
        </form>
    </td>
</tr>
