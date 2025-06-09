<script>
	import DarkToggle from "../components/DarkToggle.svelte";
	import NavButton from "../components/NavButton.svelte";

</script>
<nav class="flex p-4 items-center justify-between">
    <img class="h-16 w-auto" src="./SHELF_LOGO.svg" alt="Shelf" >
    <DarkToggle> </DarkToggle>
</nav>

<div class="flex flex-col font-serif text-center text-balance items-center justify-center h-[60svh]">
    <h1 class="text-9xl p-10" >welcome to <br> your digital <inline class="text-logo-red italic">shelf</inline></h1>
    <p class="font-sans text-lg">Digitalize your bookshelf, track progress, make a group - share your shelves.</p>
    <div class="flex p-10">
        <NavButton link="/login" buttonText="Log In" color="red"></NavButton>

        <NavButton link="/register" buttonText="Sign Up" color="blue"></NavButton>
    </div>
</div>

