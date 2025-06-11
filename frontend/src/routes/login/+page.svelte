<script lang="ts">
	import { page } from '$app/state';
	import LoginForm from '../../components/login/LoginForm.svelte';
	import SignUpForm from '../../components/login/SignUpForm.svelte';

	let isLoginSelected = page.state;
	function toggleActive() {
		isLoginSelected = !isLoginSelected;
	}
</script>

{#snippet nameField()}
	<div>
		<label for="username" class="block p-2">Username </label>
		<input name="username" type="text" />
	</div>
{/snippet}
{#snippet emailField()}
	<div>
		<label for="email" class="block p-2">Email</label>
		<input name="email" type="email" autocomplete="email" required />
	</div>
{/snippet}
{#snippet passwordField()}
	<div>
		<label for="password" class="block p-2">Password</label>
		<input name="password" type="password" autocomplete="current-password" required />
	</div>
{/snippet}

<div class="flex h-lvh flex-wrap items-center justify-center font-medium">
	<SignUpForm isSignUpSelected={!isLoginSelected} toggleActive={toggleActive} emailField={emailField} passwordField={passwordField} usernameField={nameField}/>
	<LoginForm isLoginSelected={isLoginSelected} toggleActive={toggleActive} emailField={emailField} passwordField={passwordField} />
</div>

<style>
	input {
		border: 1px solid var(--color-neutral-500);
		border-radius: var(--radius-xl);
		padding: var(--spacing);
		width: 100%;
	}
</style>
