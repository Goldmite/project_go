<script lang="ts">
	let { children } = $props();
</script>

<div
	class="my-4 grid grid-cols-[repeat(auto-fit,minmax(10.5rem,auto))] justify-start gap-6
    sm:grid-cols-[repeat(auto-fit,minmax(14rem,auto))]
	sm:gap-10"
>
	{@render children()}
</div>
