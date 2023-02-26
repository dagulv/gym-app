<script context="module">

    const modal = writable({
        component: null,
    });

    export function openModal(component) {
        return modal.set({
            component: component,
        });
    }

    function closeModal() {
        modal.set({
            component: null,
        });
    }
</script>

<script>
    import { writable } from "svelte/store";
    import Close from "$iconComponents/Close.svelte";
    
</script>

{#if $modal.component}
    <div class="overlay" on:click={closeModal}></div>
    <div class="modal">
        <button on:click={closeModal}><Close /></button>
        <svelte:component this={$modal.component} on:close={closeModal} />
    </div>
{/if}

<style lang="scss">
    .overlay {
        position: absolute;
        inset: 0;
        background-color: rgba(#000, 0.7);
        z-index: 3;
    }
    .modal {
        position: absolute;
        top: 50%;
        left: 50%;

        border-radius: $site-space * 0.5;
        transform: translate(-50%, -50%);
        max-width: 95%;
        padding: $site-space * 0.5;
        background-color: mix($color-background, #fff, 90%);
        z-index: 4;
    }
</style>