// @ts-nocheck
import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './[slug]/$types';
import { getNewEvent } from '$lib/utils/utils';
export const load = async ({ parent }: Parameters<PageServerLoad>[0]) => {
	const { token, userRole, currentUserID } = await parent();
	const event = await getNewEvent();
	if (userRole !== 'admin' && userRole !== 'manager') {
		redirect(302, '/not-allowed');
	}

	return {
		token,
		userRole,
		currentUserID,
		event
	};
};
