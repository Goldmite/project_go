<script lang="ts">
    let {link, buttonText, color} = $props();
</script>
<a href="{link}">
    <button class="text-logo-{color} font-sans bg-dark dark:bg-light p-2 font-bold rounded-xl m-4"> {buttonText} </button>
</a>