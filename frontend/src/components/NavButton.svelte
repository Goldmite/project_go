<script lang="ts">
	import { goto } from "$app/navigation";

	let { link, buttonText, color } = $props();
</script>

<!-- Safelist: text-logo-blue shadow-logo-red shadow-logo-blue -->
<button onclick={() => goto("/login", {
	state: link=="/login"
})}
	class="text-{color} bg-neutral rounded-xl font-sans font-bold shadow-md hover:shadow-lg shadow-{color} m-2 p-3"
>
	{buttonText}
</button>

