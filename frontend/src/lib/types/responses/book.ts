export interface Book {
	owned_by: string[];
	isbn: string;
	title: string;
	authors: string[];
	pages: number;
	description: string;
	publisher: string;
	publishedDate: string;
	language: string;
	cover: string;
}

export interface BookProgress {
	pages_read: number;
	time_read: number;
	first_page: number;
	current_page: number;
}

export interface RecentlyRead {
	isbn: string;
	title: string;
	authors: string[];
	cover: string;
	start_date: string;
	last_read_date: string;
}
