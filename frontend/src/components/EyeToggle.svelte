<script lang="ts">
    let { toggle = $bindable() } = $props();

</script>
<!-- svelte-ignore a11y_no_static_element_interactions -->
<!-- svelte-ignore a11y_click_events_have_key_events -->
<span
    class="{toggle
        ? 'icon-[solar--eye-closed-linear]'
        : 'icon-[solar--eye-linear]'} -ml-8 text-xl"
    onclick={() => (toggle = !toggle)}
>
</span>