<script lang="ts">
	let { children, shadowColorClass, isActive, handleClick } = $props();
</script>

<!-- svelte-ignore a11y_no_static_element_interactions -->
<!-- Safelist: login signup-->
<!-- svelte-ignore a11y_click_events_have_key_events -->
<div
	class="card {shadowColorClass} text-dark m-8 h-[500px] w-full p-10 sm:w-1/4 {!isActive &&
		'dots brightness-95 backdrop-brightness-95'}"
	class:alternate={shadowColorClass == 'signup'}
	class:active={isActive}
	onclick={isActive ? undefined : handleClick}
>
	{@render children()}
</div>

<style>
	@layer components {
		.login {
			--tw-shadow-color: var(--color-logo-red);
			--tw-inset-shadow-color: var(--color-logo-red);
		}
		.signup {
			--tw-shadow-color: var(--color-logo-blue);
			--tw-inset-shadow-color: var(--color-logo-blue);
		}
		.card {
			border-radius: var(--radius-2xl);
			border-width: 1px;
			--tw-shadow: inset 2px 10px 6px var(--tw-inset-shadow-color);
			box-shadow: var(--tw-shadow);
		}

		.card.active {
			--tw-shadow: 2px 8px 8px var(--tw-shadow-color);
			background-color: var(--color-light);
		}

		.dots {
			background-image: radial-gradient(var(--color-logo-red) 0px, transparent 2px);
			background-size: 20px 20px;
		}
		.dots.alternate {
			--color-logo-red: var(--color-logo-blue);
		}
	}
</style>
