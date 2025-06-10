<script lang="ts">
	import LoginCard from "../../components/LoginCard.svelte";

    let isSelected = $state(true);
    function toggleActive() {
        isSelected = !isSelected;
    }
</script>

<div class="flex justify-center items-center h-svh">
    <LoginCard shadowColorClass="signup" isActive={isSelected} handleClick={toggleActive}>
        <h1>  SIGN UP </h1>
    </LoginCard>

    <LoginCard shadowColorClass="login" isActive={!isSelected} handleClick={toggleActive}>
        <h1> LOGIN </h1>
    </LoginCard>
</div>
