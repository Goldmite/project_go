<script lang="ts">
	import { goto } from '$app/navigation';

	let { backUrl } = $props();
</script>

<button class="text-logo-red mr-4 font-sans" onclick={() => goto(backUrl)}>{'<-I'}</button>
