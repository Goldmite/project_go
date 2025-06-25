<script lang="ts">
	import CurrentlyReading from '../../../components/dashboard/CurrentlyReading.svelte';
	import Grid from '../../../components/dashboard/Grid.svelte';
	import NavCard from '../../../components/dashboard/NavCard.svelte';
</script>

{#snippet header(title: string)}
	<h1 class="text-lg font-bold sm:text-2xl">{title}</h1>
{/snippet}

<!--
<Grid>
    <div class="col-span-3 col-end-3">
        <h1>PERSONAL</h1>
        <Grid>
            <div class="col-span-3">
                <Grid>
                    <div class="cardContainer col-span-1">
                        <NavCard>
                            {@render CardName("Personal")}
                        </NavCard>
                    </div>
                    <div class="stats col-start-2 col-span-2">
                        <NavCard>
                            Stats or something
                        </NavCard>
                    </div>
                </Grid>
            </div>
            <div class="col-span-3 row-start-2">
                <h1>GROUPS</h1>
                <Grid>
                    <div class="cardContainer col-span-1">
                        <NavCard>
                            Group 1
                        </NavCard>
                    </div>
                    <div class="cardContainer col-span-1 col-start-2">
                        <NavCard>
                            Group 2
                        </NavCard>
                    </div>
                    <div class="cardContainer col-span-1 col-start-3">
                        <NavCard>
                            Group 3
                        </NavCard>
                    </div>
                </Grid>
            </div>
        </Grid>
    </div>
    <div class="col-start-3">
        <h1>CURRENTLY READING</h1>
        <Grid>
            <div class="cardContainer">
                <NavCard>
                    Currently reading 1
                </NavCard>
            </div>
            <div class="cardContainer">
                <NavCard>
                    Currently reading 2
                </NavCard>
            </div>
            <div class="cardContainer">
                <NavCard>
                    Currently reading 3
                </NavCard>
            </div>
        </Grid>
    </div>
    
</Grid>
-->
<div class="flex justify-between">
	<div class="flex-1">
		{@render header('Shelves')}
		<Grid>
			<NavCard>{@render header('Personal')}</NavCard>
			<NavCard>{@render header('Group1')}</NavCard>
			<NavCard>{@render header('Group2')}</NavCard>
			<NavCard>{@render header('Group3')}</NavCard>
		</Grid>
	</div>

	<CurrentlyReading>{@render header('Currently Reading')}</CurrentlyReading>
</div>

<style>
</style>
