<script lang="ts">
	import type { PageProps } from './$types';
	import PageHeader from '../../../../components/PageHeader.svelte';
	import BookGrid from '../../../../components/shelf/BookGrid.svelte';
	import MembersList from '../../../../components/shelf/shared/MembersList.svelte';
	import { preloadData, pushState } from '$app/navigation';
	import { page } from '$app/state';

	let { data }: PageProps = $props();
</script>

<span class="flex flex-col md:flex-row justify-between items-baseline">
	<PageHeader>{data.currGroup?.name} shelf</PageHeader>
	<button class="rounded-2xl min-w-28 h-12 text-dark text-4xl font-light hover:font-normal 
	active:font-normal italic bg-logo-blue hover:outline hover:shadow-logo-blue shadow-lg"
	onclick={async () => {
		const result = await preloadData(`/shared/${data.currGroup.id}/invite`);
		if (result.type === 'loaded' && result.status === 200) {
			pushState(`/shared/${data.currGroup.id}/invite`, {
				...page.state,
				invite: result.data,
				groupMembers: data.members
			}); } }}
	>+</button>
</span>
<MembersList members={data.members} />
<BookGrid books={data.books} />
