import type { PageServerLoad } from './$types';
export const load: PageServerLoad = async ({ cookies }) => {
	console.log('Removing token cookie');
	cookies.delete('token', { path: '/' });
};
