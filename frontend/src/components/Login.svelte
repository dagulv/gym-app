<script>
    import { invalidateAll } from '$app/navigation';

    let name;
    let password;
    let submitButton = 'Logga in';
    let isLogin = true;

    let url = 'http://localhost:8080/api/login';
    function formSwitch() {
        if (!isLogin) {
            submitButton = 'Logga in';
            url = 'http://localhost:8080/api/login';
            isLogin = true;
        } else {
            submitButton = 'Skapa konto';
            url = 'http://localhost:8080/api/signup';
            isLogin = false;
        }
    }

    async function onSubmit() {
        try {
            const body = {
                name: name,
                password: password,
            }

            const res = await fetch(url, {
                method: 'POST',
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(body)
            });
            const data = res.json()
            if (data) {
                console.log(url);
            }
            await invalidateAll();
        } catch (e) {
            console.log(e);
        }
    }
</script>

<form on:submit|preventDefault={onSubmit}>
    <input bind:value={name} type="text" placeholder="Name" required>
    <input bind:value={password} type="password" placeholder="Password" required>
    <button type="submit">{submitButton}</button>
</form>
<span on:click={formSwitch}>Skapa konto</span>