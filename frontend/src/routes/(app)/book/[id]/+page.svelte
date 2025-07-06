<script lang="ts">
	import { read } from '$app/server';
	import type { Book } from '$lib/types/book';
	import { toReadableLanguage } from '$lib/types/enums/language';
	import PageSubheader from '../../../../components/PageSubheader.svelte';
	import CoverWrapper from '../../../../components/shelf/CoverWrapper.svelte';
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
	const book: Book = data.book;
	let readMore = $state(false);
	const descriptionPreview = 300;
</script>

<div class="gap-2 sm:gap-6">
	<div class="float-left mt-3 mr-3 min-w-42 sm:min-w-56">
		<CoverWrapper cover={book.cover} title={book.title} />
	</div>
	<div>
		<span class="flex flex-col items-baseline sm:flex-row sm:gap-4">
			<PageSubheader>{book.title}</PageSubheader>
			<p class="inline font-light italic">
				by
				{#each book.authors as author, i}
					{author}{i + 1 < book.authors.length ? ', ' : ''}
				{/each}
			</p>
		</span>

		<p class="my-4 text-justify hyphens-auto">
			{#if readMore}
				{book.description}
			{:else}
				{book.description.slice(0, descriptionPreview)}...
			{/if}
			<button class="text-sm italic" onclick={() => (readMore = !readMore)}>
				{#if book.description.length > descriptionPreview}
					{#if readMore}
						See less
					{:else}
						Read more
					{/if}
				{/if}
			</button>
		</p>

		<div class="text-center">
			<p>{book.isbn}</p>
			<p>{book.pages} pages</p>
			<p>{toReadableLanguage(book.language)}</p>
			<p>By {book.publisher} in {book.publishedDate}</p>
		</div>
	</div>
</div>
