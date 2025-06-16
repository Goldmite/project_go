<script lang="ts">
    let { children } = $props();

</script>

<div class="grid justify-start gap-6 sm:gap-10 my-4
    grid-cols-[repeat(auto-fit,minmax(10.5rem,auto))]
	sm:grid-cols-[repeat(auto-fit,minmax(14rem,auto))]"
>
    {@render children()}
</div>