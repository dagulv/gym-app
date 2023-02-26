import { writable } from 'svelte/store';

export const sessionStore = writable({
    id: null,
    name: null,
    loggedIn: false,
});