<script lang="ts">
    let { currStep } = $props();

</script>

<div class="flex justify-between my-4">
    {#each [1, 2, 3] as step}
        <div class="flex flex-col flex-1 relative items-center text-sm">
            <div class="w-6 h-6 z-30 relative rounded-full text-center justify-center border transition-colors duration-300 {step <= currStep ? 'bg-status-logo-done text-dark dark:border-light' : 'bg-light dark:bg-dark'}">
                {step}
            </div>
            {#if step < 3}
                <div class="absolute w-full z-20 left-1/2 top-3 border"></div>
            {/if}
        </div>
    {/each}
</div>