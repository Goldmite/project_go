<script lang="ts">
    let { children, shadowColorClass, isActive, handleClick } = $props();

    
</script>
<!-- Safelist: login signup -->
<button class=" card {shadowColorClass} w-1/4 h-1/2 p-10 m-8" class:active={isActive} onclick={handleClick} disabled={isActive}>
    {@render children()}
</button>

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
            border-radius: var(--radius-xl);
            --tw-shadow: inset 0 0 20px var(--tw-shadow-color);
            box-shadow: var(--tw-shadow);
        }

        .active {
            --tw-shadow: 0 0 20px var(--tw-inset-shadow-color);
        }
    }

</style>