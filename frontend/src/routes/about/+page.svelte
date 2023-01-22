<script context="module">
    const status = fetch("http://localhost:8080/hello.json")
    .then((r) => r.json())
    .catch((err) => new Error(err));
</script>

<svelte:head>
    <title>Om</title>
</svelte:head>
<h1>About</h1>
<p>This is the about page</p>

<p>Server status:
{#await status}
    ...loading
{:then data} 
    {data.message}
{:catch err}
    failed to load data {err}
{/await}
</p>