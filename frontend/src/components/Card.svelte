<script>
    import { goto } from '$app/navigation';

    export let card = {};

    console.log(card);
    function onClick() {
        goto(`/schedules/${card.id}`);
    }
</script>
<article on:click={onClick}>
    <h2><a href={`/schedules/${card.id}`}>{card.title}</a></h2>
    <p>{card.description}</p>
</article>

<style lang="scss">
    article {
        background-color: green;
        aspect-ratio: 1/1;
        width: calc(50% - $site-space * 2);
        max-width: $site-space * 8;
        color: #fff;
        border-radius: $site-space;
        padding: $site-space;
        cursor: pointer;
        transition: $transition-time;

        &:hover {
            scale: 1.05;
        }
    }
</style>
