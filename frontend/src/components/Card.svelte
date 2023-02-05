<script>

export let card;
console.log(card);
</script>
<article>
    <!-- <h2><a href={card.id}>{card.title}</a></h2>
    <p>{card.description}</p> -->
</article>

<style lang="scss">
    article {
        background-color: green;
        aspect-ratio: 1/1;
        width: calc(50% - $site-space * 2);
        color: #fff;
        border-radius: $site-space;
        padding: $site-space;
    }
</style>
