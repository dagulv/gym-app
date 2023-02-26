<script>
    import Card from "$components/Card.svelte";
    import Add from "$iconComponents/Add.svelte";
    import { openModal } from "$components/Modal.svelte";
    import FormSchedule from "$components/FormSchedule.svelte";

    export let data;
    
    function addSchedule() {
        openModal(FormSchedule);
    }
</script>

<section>
    {#if data.schedules.length}
        {#each data.schedules as schedule}
            <Card card={schedule}/>
            <Card card={schedule}/>
        {/each}
    {:else}
        <button on:click={addSchedule}><Add /></button>
    {/if}
</section>

<style lang="scss">
    section {
        display: flex;
        gap: $site-space;
    }
</style>