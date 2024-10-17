// @ts-nocheck
import type { PageServerLoad } from './$types';
export const load = async ({ cookies }: Parameters<PageServerLoad>[0]) => {
	const token = cookies.get('token');
	return {
		token: token
	};
};
