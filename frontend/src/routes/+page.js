/** @type {import('./$types').PageLoad} */
export async function load({ fetch }) {
    const data = fetch("http://localhost:8080/schema")
        .then((r) => r.json())
        .catch((err) => new Error(err));

    return { data }
}