<script lang="ts">
	import LoginCard from './LoginCard.svelte';

	let { isSignUpSelected, toggleActive, emailField, passwordField, usernameField } = $props();
</script>

<LoginCard shadowColorClass="signup" isActive={isSignUpSelected} handleClick={toggleActive}>
	{#if isSignUpSelected}
		<form class="flex h-full flex-col space-y-6">
			{@render usernameField()}
			{@render emailField()}
			{@render passwordField()}
			<button
				class="bg-logo-blue mt-auto w-full rounded-2xl p-2 transition delay-100 duration-300 ease-in-out hover:scale-110 focus:outline-1"
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
