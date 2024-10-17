export interface Event {
	id: number;
	title: string;
	short_description: string;
	description: string;
	location: string;
	date: string;
	time: string;
	is_featured: boolean;
	created_by: number;
}

export interface EventInput {
	title: string;
	short_description: string;
	description: string;
	location: string;
	date: string;
	time: string;
	is_featured: boolean;
}
