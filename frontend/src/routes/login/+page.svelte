<script lang="ts">
	import { page } from '$app/state';
	import LoginForm from '../../components/login/LoginForm.svelte';
	import SignUpForm from '../../components/login/SignUpForm.svelte';
	import { Tween } from 'svelte/motion';
	import { zxcvbn } from '@zxcvbn-ts/core';
	import { cubicOut } from 'svelte/easing';
	import EyeToggle from '../../components/EyeToggle.svelte';
	import type { PageProps } from './$types';
	import ErrorMsg from '../../components/errors/ErrorMsg.svelte';

	let { form }: PageProps = $props();
	let isLoginSelected = $derived(page.state.isLoginSelected);
	function toggleActive() {
		isLoginSelected = !isLoginSelected;
	}
	$effect(() => {
		if (form?.success) {
			// switch to login form after sign up form success
			isLoginSelected = true;
		}
	});
	let showPassword = $state(false);
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
		<label for="name" class="block p-2">Username </label>
		<input
			class="outline-status-logo-done focus:invalid:outline-logo-red focus:outline-3"
			name="name"
			type="text"
			minlength="2"
			maxlength="20"
			required
		/>
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
			maxlength="254"
			required
		/>
	</div>
{/snippet}

{#snippet passwordField()}
	<div>
		<div>
			<label for="password" class="block p-2">Password</label>
			<div class="flex items-center">
				<input
					class="outline-status-logo-done focus:invalid:outline-logo-red focus:outline-3"
					bind:value={password}
					name="password"
					type={showPassword ? 'text' : 'password'}
					maxlength="128"
					required
					bind:this={inputStrength}
				/>
				<EyeToggle bind:toggle={showPassword} />
			</div>
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
	>
		{#if form?.emailtaken}
			<div class="absolute top-0 w-full text-center">
				<ErrorMsg msg="| EMAIL TAKEN |" />
			</div>
		{/if}
	</SignUpForm>
	<LoginForm {isLoginSelected} {toggleActive} {emailField} {passwordField}>
		{#if form?.notfound}
			<div class="absolute top-0 w-full text-center">
				<ErrorMsg msg="| INCORRECT CREDENTIALS |" />
			</div>
		{/if}
	</LoginForm>
</div>

<style>
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
