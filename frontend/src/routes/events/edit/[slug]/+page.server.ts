import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import type { User } from '$lib/schemas/user';
import { getEvent, getUser } from '$lib/utils/utils';
import type { Event } from '$lib/schemas/event';
export const load: PageServerLoad = async ({ params, cookies }) => {
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

	const event: Event = await getEvent(Number(params.slug), token);
	if (event.created_by !== user.id && user.role !== 'admin') {
		redirect(302, '/not-allowed');
	}

	return {
		token: token,
		userRole: user.role,
		currentUserID: user.id,
		event: event
	};
};
