import { sessionStore } from '$lib/stores';
/** @type {import('./$types').LayoutLoad} */
export async function load({ fetch }) {
    let session;
    try {
        const resp = await fetch("http://localhost:8080/api/users/me", {
            headers: {
                'Accept': 'application/json',
            },
            credentials: 'include',
        });
        session = await resp.json();

        if (resp.status === 200) {
            sessionStore.set({
                id: session.id,
                name: session.name,
                loggedIn: true,
            });
        }
    } catch (e) {
        console.log(e);
    }
}

export const prerender = false;
export const ssr = false;