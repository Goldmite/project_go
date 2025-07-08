<script lang="ts">
	import BackButton from '../../../../../components/navigation/BackButton.svelte';

    let { children } = $props();

</script>

<div class="m-4">
<span class="absolute font-bold text-2xl mt-6">
    <BackButton />
</span>
<div class="flex justify-center items-center h-screen max-h-screen">
    {@render children()}
</div>
</div>