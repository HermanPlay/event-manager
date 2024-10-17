import { writable } from 'svelte/store';

export const toasts = writable([] as Toast[]);

interface Toast {
	id: number;
	type: string;
	message: string;
	dismissible: boolean;
	timeout: number;
}

export const addToast = (toast: {
	type: string;
	message: string;
	dismissible?: boolean;
	timeout?: number;
}) => {
	// Create a unique ID so we can easily find/remove it
	// if it is dismissible/has a timeout.
	const id = Math.floor(Date.now());

	// Setup some sensible defaults for a toast.
	const defaults: {
		id: number;
		type: string;
		dismissible: boolean;
		timeout: number;
	} = {
		id,
		type: 'info',
		dismissible: true,
		timeout: 3000
	};

	// Push the toast to the top of the list of toasts
	toasts.update((all: Toast[]) => [{ ...defaults, ...toast }, ...all]);

	// If toast is dismissible, dismiss it after "timeout" amount of time.
	if (toast.timeout) setTimeout(() => dismissToast(id), toast.timeout);
};

export const dismissToast = (id: number) => {
	toasts.update((all) => all.filter((t) => t.id !== id));
};
