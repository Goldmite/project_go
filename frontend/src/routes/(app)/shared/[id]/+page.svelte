<script lang="ts">
	import type { PageProps } from './$types';
	import PageHeader from '../../../../components/PageHeader.svelte';
	import BookGrid from '../../../../components/shelf/BookGrid.svelte';
	import MembersList from '../../../../components/shelf/shared/MembersList.svelte';
	import { preloadData, pushState } from '$app/navigation';
	import { page } from '$app/state';

	let { data }: PageProps = $props();
</script>

<PageHeader>{data.currGroup?.name} shelf</PageHeader>
<MembersList members={data.members}>
	<button class="border p-2" 
			onclick={async () => {
				const result = await preloadData(`/shared/${data.currGroup.id}/invite`);
				if (result.type === 'loaded' && result.status === 200) {
					pushState(`/shared/${data.currGroup.id}/invite`, {
						...page.state,
						invite: result.data
					});
				}
			}}
	>Invite</button>
</MembersList>
<BookGrid books={data.books} />
