import { persisted } from 'svelte-persisted-store';

export const isAuthenticated = persisted('isAuthenticated', false);
