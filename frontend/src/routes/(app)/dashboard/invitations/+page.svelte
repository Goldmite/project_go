<script lang="ts">
	import InvitationRow from "../../../../components/dashboard/invitations/InvitationRow.svelte";
	import type { PageProps } from "./$types";

    let { data }: PageProps = $props();

</script>

<div class="overflow-y-scroll h-full px-1 bg-light text-dark">
    <table class="table-auto">
        <thead class="bg-light sticky top-0 z-10">
            <tr>
                <th>Group</th>
                <th>Invited by</th>
                <th></th>
            </tr>
        </thead>
        <tbody>
            {#each data.invitations as inv}
                <InvitationRow {inv} />
            {/each}
        </tbody>
    </table>
</div>
