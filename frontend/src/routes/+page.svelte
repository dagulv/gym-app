<script>
  /** @type {import('./$types').PageData} */  
  export let data;
</script>
<svelte:head>
    <title>Homepage</title>
</svelte:head>
<h1>WelcomeSvelteKit</h1>
<p>{console.log(data.status.title)}</p>