<script lang="ts">
	import LoginCard from "../../components/LoginCard.svelte";

    let isLoginSelected = $state(true);
    function toggleActive() {
        isLoginSelected = !isLoginSelected;
    }
</script>
{#snippet emailField()}
    <div>
        <label for="email" class="block p-2">Email</label>
        <input name="email" type="email" autocomplete="email" required>
    </div>
{/snippet}
{#snippet passwordField()}
    <div>
        <label for="password" class="block p-2">Password</label>
        <input name="password" type="password" autocomplete="current-password" required>
    </div>
{/snippet}

<div class="flex flex-wrap justify-center items-center h-lvh font-medium">
    <LoginCard shadowColorClass="signup" isActive={!isLoginSelected} handleClick={toggleActive}>
        {#if !isLoginSelected }
        <form class="space-y-6 h-full flex flex-col">
            <div>
                <label for="username" class="block p-2">Username </label>
                <input name="username" type="text">
            </div>
            {@render emailField()}
            {@render passwordField()}
            <button class="w-full bg-logo-blue rounded-2xl p-2 mt-auto">
                Sign up
            </button>
        </form>
        {:else}
            <div class="flex items-center h-full font-serif italic font-light tracking-wide">
            <h1 class="w-full text-center text-8xl/30 text-logo-blue/80 dark:text-logo-blue/60 text-shadow-lg/15" >Sign<br>up</h1>
        </div>
        {/if}
    </LoginCard>

    <LoginCard shadowColorClass="login" isActive={isLoginSelected} handleClick={toggleActive}>
        {#if isLoginSelected }
        <form class="space-y-6 h-full flex flex-col">
            <div class="w-full space-y-6 flex-1">
            {@render emailField()}
            {@render passwordField()}
            </div>
            <button class="w-full bg-logo-red rounded-2xl p-2 mt-auto">
                Log in
            </button>
        </form>
        {:else}
        <div class="flex items-center h-full font-serif italic font-light tracking-wide">
            <h1 class="w-full text-center text-8xl/30 text-logo-red/80 dark:text-logo-red/60 text-shadow-lg/15" >Log<br>in</h1>
        </div>
        {/if}
    </LoginCard>
</div>

<style>
    input {
        border: 1px solid var(--color-neutral-500);
        border-radius: var(--radius-xl);
        padding: var(--spacing);
        width: 100%;
    }
</style>
