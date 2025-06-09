<script lang="ts">
    let {link, buttonText, color} = $props();
</script>
<a href="{link}">
    <!-- Safelist: text-logo-blue shadow-logo-red shadow-logo-blue -->
    <button class="text-{color} bg-neutral font-sans font-bold rounded-xl shadow-md hover:shadow-lg shadow-{color} p-3 m-2 "> {buttonText} </button>
</a>