import type { PageServerLoad } from './$types';
import { getUser } from '$lib/utils/utils';
import { redirect } from '@sveltejs/kit';
import type { User } from '$lib/schemas/user';
export const load: PageServerLoad = async ({ cookies }) => {
	const token = cookies.get('token');
	if (!token) {
		throw redirect(302, '/auth');
	}
	let user: User;
	try {
		user = await getUser(token);
	} catch (error: any) {
		console.error(error);
		return { error: error.message };
	}
	if (user.role !== 'admin') {
		throw redirect(302, '/not-allowed');
	}

	return {
		token: token,
		userRole: user.role,
		currentUserID: user.id
	};
};
