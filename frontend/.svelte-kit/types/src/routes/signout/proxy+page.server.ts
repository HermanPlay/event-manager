// @ts-nocheck
import type { PageServerLoad } from './$types';
export const load = async ({ cookies }: Parameters<PageServerLoad>[0]) => {
	console.log('Removing token cookie');
	cookies.delete('token', { path: '/' });
};
