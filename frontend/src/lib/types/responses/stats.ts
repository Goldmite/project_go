export interface Stats {
	total_pages: number;
	total_time: number;
}

export interface ReadingSession {
	date: string;
	time_read: number;
}

export type Heatmap = {
	[week: number]: {
		[day: number]: {
			date: string;
			time: number;
			level: number;
		};
	};
};
