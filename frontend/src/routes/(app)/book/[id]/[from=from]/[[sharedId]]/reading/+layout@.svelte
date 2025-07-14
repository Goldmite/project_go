<script lang="ts">
	import { page } from '$app/state';
	import BackButton from '../../../../../../../components/navigation/BackButton.svelte';

	let { children } = $props();
</script>

<div class="m-4">
	<span class="absolute mt-6 text-2xl font-bold">
		<BackButton backUrl="/book/{page.params.id}/{page.params.from}/{page.params.sharedId}" />
	</span>
	<div class="flex h-screen max-h-screen items-center justify-center">
		{@render children()}
	</div>
</div>
