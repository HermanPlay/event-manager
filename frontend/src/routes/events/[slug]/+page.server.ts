import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import { getEvent } from '$lib/utils/utils';

export const load: PageServerLoad = async ({ parent, params }) => {
	const { token, userRole, currentUserID } = await parent();
	if (!token) {
		throw redirect(302, '/auth');
	}
	const event = await getEvent(Number(params.slug), token);

	return {
		token,
		userRole,
		currentUserID,
		event
	};
};
