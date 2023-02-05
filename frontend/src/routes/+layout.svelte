<script>
    import Header from "../components/Header.svelte";
</script>

<Header />

<main>
    <slot />
</main>

<style lang="scss">
    :global {
        @import '../styles/frontend.scss';
    }
    :global(body) {
        background-color: $color-background;
    }
    main {
        padding: $site-space;
    }
    
</style>