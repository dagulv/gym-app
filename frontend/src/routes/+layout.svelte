<script>
    import Header from "$components/Header.svelte";
    import Login from "$components/Login.svelte";
    import { sessionStore } from "$lib/stores.js"
    import Modal from "$components/Modal.svelte";
</script>

{#if $sessionStore.loggedIn}
<Header />
<Modal />
<main>
    <slot />
</main>
{:else}
<main>
    <Login />
</main>
{/if}

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