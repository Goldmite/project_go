<script lang="ts">
    let { showModal = $bindable(), header, children } = $props();
    let dialog: HTMLDialogElement | undefined = $state();

    $effect(() => {
        if (showModal) dialog?.showModal();
    });
</script>

<dialog
    bind:this={dialog}
    onclose={() => (showModal = false)}
    onclick={(e) => { if (e.target === dialog) dialog.close(); }}
>
    <div>
        {@render header()}
        {@render children()}
        <button onclick={() => dialog?.close()}>Close</button>
    </div>
</dialog>