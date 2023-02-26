<script>
    import { sessionStore } from "$lib/stores.js"
    import { createEventDispatcher } from "svelte";

    let title;
    let description;

    async function onSubmit() {
        try {
            const body = {
                userId: sessionStore.id,
                title: title,
                description: description,
            }

            const res = await fetch("http://localhost:8080/api/schedules", {
                method: 'POST',
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(body)
            });
            const data = res.json();
            if(data) {
                console.log(data);
                dispatch('close');
            }
        } catch (error) {
            console.log(error);
        }
    }

    const dispatch = createEventDispatcher();
</script>

<section>
    <h2>Skapa schema</h2>
    <form on:submit={onSubmit}>
        <div class="form-row">
            <label>
                Titel
                <input type="text" bind:value={title}>
            </label>
        </div>
        <div class="form-row">
            <label>
                Beskrivning
                <input type="text" bind:value={description}> 
            </label>
        </div>
        <div class="form-row">
            <button type="submit">Skapa schema</button>
        </div>
    </form>
</section>