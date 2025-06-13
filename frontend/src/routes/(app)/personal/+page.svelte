<script lang="ts">
	import BookGrid from "../../../components/shelf/BookGrid.svelte";
    export let data;
    
</script>
<div class="m-20 p-10">
    <h1 class="text-4xl font-bold p-4">Your shelf</h1>
    <BookGrid books={data}/>
</div>