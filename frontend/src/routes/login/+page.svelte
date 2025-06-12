<script lang="ts">
	import { page } from '$app/state';
	import LoginForm from '../../components/login/LoginForm.svelte';
	import SignUpForm from '../../components/login/SignUpForm.svelte';
	import { Tween } from 'svelte/motion';
	import { zxcvbn } from '@zxcvbn-ts/core';
	import { cubicOut } from 'svelte/easing';

	let isLoginSelected = $derived(page.state);
	function toggleActive() {
		isLoginSelected = !isLoginSelected;
	}

	let password = $state('');
	let email = $state('');

	const { score } = $derived(zxcvbn(password));
	const strength = new Tween(0, {
		duration: 500,
		easing: cubicOut
	});
	$effect(() => {
		strength.set(score);
	});

	let inputStrength: any = $state();
	$effect(() => {
		if (inputStrength && !isLoginSelected) {
			if (score < 3) {
				inputStrength.setCustomValidity('Password is too weak.');
			} else {
				inputStrength.setCustomValidity('');
			}
		}
	});
</script>

{#snippet nameField()}
	<div>
		<label for="username" class="block p-2">Username </label>
		<input class="outline-status-logo-done focus:outline-3" name="username" type="text" />
	</div>
{/snippet}

{#snippet emailField()}
	<div>
		<label for="email" class="block p-2">Email</label>
		<input
			class="outline-status-logo-done focus:invalid:outline-logo-red focus:outline-3"
			bind:value={email}
			name="email"
			type="email"
			required
		/>
	</div>
{/snippet}

{#snippet passwordField()}
	<div>
		<div>
			<label for="password" class="block p-2">Password</label>
			<input
				class="outline-status-logo-done focus:invalid:outline-logo-red focus:outline-3"
				bind:value={password}
				name="password"
				type="password"
				required
				bind:this={inputStrength}
			/>
		</div>
		{#if !isLoginSelected}
			<progress
				value={strength.current}
				max="4"
				class="bg-light h-2"
				style:--bar-color|important={score > 2
					? 'var(--color-status-logo-done)'
					: score == 2
						? 'var(--color-status-logo-progress)'
						: 'var(--color-logo-red)'}
			></progress>
		{/if}
	</div>
{/snippet}

<div class="flex h-lvh flex-wrap items-center justify-center font-medium">
	<SignUpForm
		isSignUpSelected={!isLoginSelected}
		{toggleActive}
		{emailField}
		{passwordField}
		usernameField={nameField}
	/>
	<LoginForm {isLoginSelected} {toggleActive} {emailField} {passwordField} />
</div>

<style>
	input {
		border: 1px solid var(--color-neutral-500);
		border-radius: var(--radius-xl);
		padding: var(--spacing);
		width: 100%;
	}
	progress {
		width: 100%;
		border-radius: var(--radius-xl);
		accent-color: var(--bar-color);
	}
	progress::-webkit-progress-value,
	progress::-moz-progress-bar {
		background-color: var(--bar-color);
	}
</style>
