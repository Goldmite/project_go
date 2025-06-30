<script lang="ts">
	import EyeToggle from "../../EyeToggle.svelte";
	import MembersList from "./MembersList.svelte";

    let { invites, inviteState = undefined, header, inviteToggle = $bindable(), showToggle = true } =$props();
</script>

<div
    class="inset-shadow-md dots h-full rounded-2xl border overflow-y-scroll backdrop-brightness-95
    {invites.length !== 0 && 'alternate'}"
>   
    <div class="flex justify-between items-center bg-logo-red/80 border-b border-dashed backdrop-brightness-50 px-3">
        <h1 class="text-dark/80 font-serif font-light italic text-2xl tracking-tight text-shadow-lg/15">
            {header}</h1>
        {#if showToggle}
            <EyeToggle bind:toggle={inviteToggle} />
        {/if}
    </div>
    {#if invites.length !== 0}
        <MembersList members={invites} {inviteState}></MembersList>
    {/if}
</div>