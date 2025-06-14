<script lang="ts">
    let { header = () => {}, children } = $props();
    
</script>

<div class="w-full h-full p-2 border-2">
    <div class="text-dark  h-full">
        {@render header()}
        {@render children()}
    </div>
</div>
