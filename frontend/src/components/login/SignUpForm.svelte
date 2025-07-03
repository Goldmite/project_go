<script lang="ts">
	import LoginCard from './LoginCard.svelte';
	import { enhance } from '$app/forms';

	let { children = () => {}, isSignUpSelected, toggleActive, emailField, passwordField, usernameField } = $props();
</script>

<LoginCard shadowColorClass="signup" isActive={isSignUpSelected} handleClick={toggleActive}>
	{#if isSignUpSelected}
		<form method="POST" use:enhance action="?/signup" class="relative flex h-full flex-col space-y-6">
			{@render children()}
			{@render usernameField()}
			{@render emailField()}
			{@render passwordField()}
			<button
				class="bg-logo-blue mt-auto w-full rounded-2xl p-2 transition delay-50 duration-200 ease-in-out hover:scale-110 focus:outline-1"
			>
				Sign up
			</button>
		</form>
	{:else}
		<div class="flex h-full items-center font-serif font-light tracking-wide italic">
			<h1
				class="text-logo-blue/80 dark:text-logo-blue/60 w-full text-center text-8xl/30 text-shadow-lg/15"
			>
				Sign<br />up
			</h1>
		</div>
	{/if}
</LoginCard>
