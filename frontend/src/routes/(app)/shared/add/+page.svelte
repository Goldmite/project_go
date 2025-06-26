<script lang="ts">
	import { enhance } from "$app/forms";
	import type { User } from "$lib/types/user";
	import ErrorMsg from "../../../../components/errors/ErrorMsg.svelte";
	import ModalWrapper from "../../../../components/ModalWrapper.svelte";
	import PageSubheader from "../../../../components/PageSubheader.svelte";
	import MembersList from "../../../../components/shelf/shared/MembersList.svelte";
	import StepProgressBar from "../../../../components/StepProgressBar.svelte";
	import type { PageProps } from "./$types";

    let { data, form }: PageProps = $props();
	let isOpen = $state(false);
    let step = $state(1);
    const invitees: User[] = $state([]);
    const invIds = new Set<string>();

    function handleArrows(next: boolean) {
        if (next) {
            if (step < 3) step++;
        } else {
            if (step > 1) step--;
        }
    };
    $effect(() => {
        if (form?.success && !invIds.has(form.invitee.id) && invitees.length < 10) {
            invitees.push({
                id: form.invitee.id,
                name: form.invitee.name,
                email: form.invitee.email
            });
            invIds.add(form.invitee.id);
        }
    });
</script>

{#snippet arrowButton(next: boolean)}
    <button 
        class="active:inset-shadow-md mx-4 w-full rounded-2xl p-2 text-lg hover:outline-1 focus:outline-1 active:text-base {next ? `bg-logo-blue ${step == 3 ? 'hidden' : ''}` : `bg-logo-red ${step == 1 ? 'hidden' : ''}`}"
        onclick={() => handleArrows(next)}>
        {next ? "->" : "<-"}
    </button>
{/snippet}

<ModalWrapper bind:isOpen={isOpen}>
<div class="flex flex-col font-sans bg-light text-dark border rounded-2xl h-[500px] w-full sm:w-1/4 px-8 py-4">
    <StepProgressBar currStep={step} />
    <div class="mx-4">
        {#if step == 1}
            <PageSubheader>I. Name your group</PageSubheader>
        {:else if step == 2}
            <PageSubheader>II. Invite your friends</PageSubheader>
        {:else}
            <PageSubheader>III. Create shared shelf</PageSubheader>
        {/if}
    </div>
    <form method="POST" use:enhance class="flex flex-col m-4">
        {#if step == 1}
            <input 
                class="outline-status-logo-done focus:invalid:outline-logo-red focus:outline-3"
                name="name"
                type="text"
                placeholder="Be creative..."
                maxlength="20"
                required
            >
        {:else if step == 2}
            <div class="w-full flex flex-row mb-2">
                <input 
                    class="w-4/5 outline-status-logo-done focus:invalid:outline-logo-red focus:outline-3"
                    
                    name="email"
                    type="email"
                    placeholder="Add friends..."
                    maxlength="256"
                    required
                >
                <button class="w-1/5 bold text-2xl "
                    formaction="?/checkUser"
                >+</button>
            </div>
            {#if form?.notfound}
                <ErrorMsg msg="I. User not found."></ErrorMsg>
            {/if}
            {#if invitees.length === 10}
                <ErrorMsg msg="II. Can't invite more users."></ErrorMsg>
            {/if}
            
            {#if invitees.length !== 0}
                <MembersList members={invitees}></MembersList>
            {/if}
        {:else}
            <button>Create</button>
        {/if}
    </form>
    
    <div class="flex w-full justify-around mt-auto">
        {#if step == 1}
            <button 
            class="active:inset-shadow-md mx-4 w-full rounded-2xl p-2 text-lg hover:outline-1 focus:outline-1 active:text-base bg-logo-red"
            onclick={() => history.back()}>x</button>
        {/if}
        {@render arrowButton(false)}
        {@render arrowButton(true)}
        {#if step == 3}
            <button 
            class="active:inset-shadow-md mx-4 w-full rounded-2xl p-2 text-lg hover:outline-1 focus:outline-1 active:text-base bg-logo-blue"
            onclick={() => isOpen = false}>v/</button>
        {/if}
    </div>
</div>
</ModalWrapper>