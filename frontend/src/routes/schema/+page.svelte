<script>
import Check from "$iconComponents/Check.svelte";
import Card from "$components/Card.svelte";

    export let data;

    console.log(data);

    const card = data.schedule.map(item => [
        {
            name: 'title',
            content: item.title
        },
        {
            name: 'description',
            content: item.description
        },
        {
            name: 'action',
            attr: {
                href: `/schema/${item.id}`
            }
        }
    ]);
</script>

<section>
    <ScheduleCard {card}/>
    <ScheduleCard {card}/>
</section>
<Check />

<style lang="scss">
    section {
        display: flex;
        gap: $site-space;
    }
</style>