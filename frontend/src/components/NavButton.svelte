<script lang="ts">
	let { link, buttonText, color } = $props();
</script>

<a href={link}>
	<!-- Safelist: text-logo-blue shadow-logo-red shadow-logo-blue -->
	<button
		class="text-{color} bg-neutral rounded-xl font-sans font-bold shadow-md hover:shadow-lg shadow-{color} m-2 p-3"
	>
		{buttonText}
	</button>
</a>
