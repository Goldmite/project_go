<script lang="ts">

    let { username, email } = $props();
</script>

<div class="group bg-logo-blue text-dark rounded-full h-6 px-4 mx-2 sm:mx-4">
    <span class="block group-active:hidden">{username}</span>
    <span class="hidden group-active:block">{email}</span>
</div>
