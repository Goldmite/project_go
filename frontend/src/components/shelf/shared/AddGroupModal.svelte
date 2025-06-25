<script lang="ts">
	import { enhance } from "$app/forms";
	import PageHeader from "../../PageHeader.svelte";
	import StepProgressBar from "../../StepProgressBar.svelte";
	import MembersList from "./MembersList.svelte";

    let { isOpen = $bindable() } = $props();
    let step = $state(1);
    let inputEmail = $state('');
    let emails: string[] = [];

    function handleArrows(next: boolean) {
        if (next) {
            if (step < 3) step++;
        } else {
            if (step > 1) step--;
        }
    }
    function handleInvite() {
        // get user by email to check if exists, 
        // if ok then add email 
        // otherwise show Not found or No such user
        emails.push(inputEmail)
    }
</script>

{#snippet arrowButton(next: boolean)}
    <button 
        class="active:inset-shadow-md mx-4 w-full rounded-2xl p-2 text-lg hover:outline-1 focus:outline-1 active:text-base {next ? `bg-logo-blue ${step == 3 ? 'hidden' : ''}` : `bg-logo-red ${step == 1 ? 'hidden' : ''}`}"
        onclick={() => handleArrows(next)}>
        {next ? "->" : "<-"}
    </button>
{/snippet}

<div class="flex flex-col font-sans border rounded-2xl h-[500px] w-full sm:w-1/4 px-8 py-4">
    <PageHeader>Create shared shelf</PageHeader>
    <StepProgressBar currStep={step} />
    
    <form method="POST" use:enhance class="flex flex-col m-4">
        {#if step == 1}
            <label for="name" class="block p-2">Group name</label>
            <input 
                class="outline-status-logo-done focus:invalid:outline-logo-red focus:outline-3"
                name="name"
                type="text"
                placeholder="Be creative..."
                maxlength="20"
                required
            >
        {:else if step == 2}
            <label for="name" class="block p-2">Invite others</label>
            <div class="w-full flex flex-row ">
                <input 
                    class="w-4/5 outline-status-logo-done focus:invalid:outline-logo-red focus:outline-3"
                    bind:value={inputEmail}
                    onkeydown={handleInvite}
                    name="email"
                    type="email"
                    placeholder="Add friends..."
                    maxlength="256"
                    required
                >
                <button class="w-1/5 bold text-2xl "
                    onclick={handleInvite}
                >+</button>
            </div>
            <!-- TODO: FIX TO GET USERS AND THEN PASS NAME AND EMAIL TO MEMBERS FIELD-->
            <MembersList members={emails}></MembersList>
        {:else}
            <button>Create</button>
        {/if}
    </form>
    
    <div class="flex w-full justify-around mt-auto">
        {#if step == 1}
            <button 
            class="active:inset-shadow-md mx-4 w-full rounded-2xl p-2 text-lg hover:outline-1 focus:outline-1 active:text-base bg-logo-red"
            onclick={() => isOpen = false}>x</button>
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