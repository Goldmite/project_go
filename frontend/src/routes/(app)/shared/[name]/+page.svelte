<script lang="ts">
	import type { PageProps } from "./$types";
	import PageHeader from "../../../../components/PageHeader.svelte";
	import BookGrid from "../../../../components/shelf/BookGrid.svelte";

    let { data }: PageProps = $props();

</script>

<PageHeader>{data.group?.name.split('-', 1)} shelf</PageHeader>
<BookGrid books={data.books} />