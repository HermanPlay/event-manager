import type { ApiResponse } from '$lib/schemas/apiResponse';
import type { Event, EventInput } from '$lib/schemas/event';
import type { User, UserInput } from '$lib/schemas/user';

export async function getUser(token: string): Promise<User> {
	// Get the user role from the token
	let response: Response;
	try {
		response = await fetch(`http://api:8080/api/user/decode`, {
			method: 'GET',
			headers: {
				'Content-Type': 'application/json',
				Authorization: `Bearer ${token}`
			}
		});
	} catch (error: any) {
		console.error(error);
		throw new Error('Failed to get user object from token');
	}

	if (!response.ok) {
		throw new Error('Failed to get user object from token');
	}
	const data = await response.json();
	return data.data;
}

export function getNewEvent(): EventInput {
	return {
		title: 'New Title',
		short_description: 'New Short Description',
		description: 'New Description',
		location: 'Warsaw, Poland',
		date: Date.now().toString(),
		time: '00:00',
		is_featured: false
	};
}

export function getNewUser(): UserInput {
	return {
		name: '',
		email: '',
		role: ''
	};
}

export function getEvent(eventId: number, token: string): Promise<Event> {
	return fetch(`http://api:8080/api/event/${eventId}`, {
		method: 'GET',
		headers: {
			'Content-Type': 'application',
			Authorization: `Bearer ${token}`
		}
	})
		.then((response) => response.json())
		.then((data: ApiResponse) => {
			return data.data;
		});
}
