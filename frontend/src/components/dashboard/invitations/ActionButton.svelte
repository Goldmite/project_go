<script lang="ts">
    let { toAccept, action = $bindable('') } = $props();
</script>

<button class="active:inset-shadow-md w-10 sm:w-14 rounded-2xl p-1 text-lg text-dark hover:outline-1 focus:outline-1 active:text-base {toAccept
    ? `bg-status-logo-done  ${action === '' ? '' : action === 'accept' ? 'disabled' : 'invisible'}`
    : `bg-logo-red ${action === '' ? '' : action === 'decline' ? 'disabled' : 'invisible'}`}"
    type="submit"
    onclick={() => {toAccept ? action = 'accept' : action = 'decline'}}
    formaction={toAccept ? '?/accept' : '?/decline'}
>
{#if toAccept}
    {'\u2713'}
{:else}
    {'\u2717'}
{/if}
</button>