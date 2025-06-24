<script lang="ts">
	import type { PageProps } from "./$types";
	import PageHeader from "../../../../components/PageHeader.svelte";
	import BookGrid from "../../../../components/shelf/BookGrid.svelte";
	import MembersList from "../../../../components/shelf/shared/MembersList.svelte";

    let { data }: PageProps = $props();

</script>

<PageHeader>{data.currGroup?.name.split('-', 1)} shelf</PageHeader>
<MembersList members={data.members}>
	<button class=" p-2 border">TEMPORARY INVITE BUTTON</button>
</MembersList>
<BookGrid books={data.books} />