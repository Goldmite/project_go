<script lang="ts">
	import InvitationRow from '../../../../components/dashboard/invitations/InvitationRow.svelte';
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
</script>


<div class="h-fit w-[336px] overflow-hidden rounded-2xl outline">
	<div class="-mr-0.5 h-fit max-h-50 min-h-10 overflow-y-scroll">
		<table class="h-full w-full table-auto font-sans">
			<tbody>
				{#each data.invitations as inv, i}
					<InvitationRow {inv} row={i + 1} />
				{:else}
					<tr>
						<td class="text-center h-10 text-lg italic tracking-wider"> No invitations </td>
					</tr>
				{/each}
			</tbody>
		</table>
	</div>
</div>
