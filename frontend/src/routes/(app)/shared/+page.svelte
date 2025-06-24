<script lang="ts">
	import { goto } from '$app/navigation';
	import PageHeader from '../../../components/PageHeader.svelte';
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
</script>

<!-- TEMPORARY PAGE - TODO: REMOVE AFTER BETTER NAVIGATION IS DONE -->
<PageHeader>Shared shelves</PageHeader>
{#each data.groupsData as g}
	<button onclick={() => goto(`/shared/${g.name}`)}>
		{g.name}
	</button>
{/each}
