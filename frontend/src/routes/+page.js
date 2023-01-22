/** @type {import('./$types').PageLoad} */
export async function load({ fetch }) {
    const status = fetch("http://localhost:8080/")
    .then((r) => r.json())
    .catch((err) => new Error(err));

    return { status }
}