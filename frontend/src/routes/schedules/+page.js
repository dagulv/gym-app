/** @type {import('./$types').PageLoad} */
export async function load({ fetch }) {
    let schedules;
    try {
        const resp = await fetch("http://localhost:8080/api/schedules", {
            headers: {
                'Accept': 'application/json',
            },
            credentials: 'include',
        });
        schedules = await resp.json();
    } catch (e) {
        console.log(e);
    }

    return {
        schedules: schedules
    }
}