<script lang="ts">
	import LoginCard from './LoginCard.svelte';
	import { enhance } from '$app/forms';

	let { children = () => {}, isLoginSelected, toggleActive, emailField, passwordField } = $props();
</script>

<LoginCard shadowColorClass="login" isActive={isLoginSelected} handleClick={toggleActive}>
	{#if isLoginSelected}
		<form method="POST" use:enhance action="?/login" class="relative flex h-full flex-col space-y-6">
			{@render children()}
			<div class="w-full flex-1 space-y-6">
				{@render emailField()}
				{@render passwordField()}
			</div>
			<button
				class="bg-logo-red mt-auto w-full rounded-2xl p-2 transition delay-50 duration-200 ease-in-out hover:scale-110 focus:outline-1"
			>
				Log in
			</button>
		</form>
	{:else}
		<div class="flex h-full items-center font-serif font-light tracking-wide italic">
			<h1
				class="text-logo-red/80 dark:text-logo-red/60 w-full text-center text-8xl/30 text-shadow-lg/15"
			>
				Log<br />in
			</h1>
		</div>
	{/if}
</LoginCard>
