<script lang="ts">
	import type { PageProps } from './$types';
	import PageHeader from '../../../../components/PageHeader.svelte';
	import BookGrid from '../../../../components/shelf/BookGrid.svelte';
	import MembersList from '../../../../components/shelf/shared/MembersList.svelte';
	import { preloadData, pushState } from '$app/navigation';
	import { page } from '$app/state';
	import { groupMembers } from '$lib/stores/localMembers';

	let { data }: PageProps = $props();
	$effect(() => {
		if (data.members) groupMembers.set(data.members);
	});
</script>

<span class="flex flex-col items-baseline md:flex-row">
	<PageHeader>{data.currGroup?.name} shelf</PageHeader>
	<button
		class="text-dark bg-logo-blue hover:shadow-logo-blue ml-auto h-12 min-w-28 rounded-2xl text-4xl
	font-light italic shadow-lg hover:font-normal hover:outline active:font-normal"
		onclick={async () => {
			const result = await preloadData(`/shared/${data.currGroup.id}/invite`);
			if (result.type === 'loaded' && result.status === 200) {
				pushState(`/shared/${data.currGroup.id}/invite`, {
					...page.state,
					invite: result.data,
					groupMembers: data.members
				});
			}
		}}>+</button
	>
</span>
<MembersList members={data.members} />
<BookGrid books={data.books} />
