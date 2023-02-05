<script>
    import MainNav from "./MainNav.svelte";
</script>

<header>
    <div class="header-wrapper">
        <MainNav />
    </div>
</header>

<style lang="scss">
    header {
        height: $site-space * 4;
        display: flex;
        align-items: center;
        justify-content: center;
        border-bottom: 1px solid $color-dark-400;
        
        .header-wrapper {
            width: 100%;
        }
    }
</style>