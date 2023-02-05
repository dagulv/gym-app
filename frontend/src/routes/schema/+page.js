/** @type {import('./$types').PageLoad} */
export async function load({ fetch }) {
    const schedule = fetch("http://localhost:8080/api/schedules")
        .then((r) => r.json())
        .catch((err) => new Error(err));

    return { schedule }
}