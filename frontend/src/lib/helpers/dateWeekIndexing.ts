import type { Heatmap, ReadingSession } from '$lib/types/responses/stats';

export function mapToWeekIndexing(sessions: ReadingSession[]): Heatmap {
	const today = normalizeDate(new Date());
	const start = normalizeDate(new Date(today));
	start.setUTCFullYear(today.getUTCFullYear() - 1);
	const msInDay = 1000 * 60 * 60 * 24;
	const daysInterval = Math.floor((today.getTime() - start.getTime()) / msInDay) + 1;
	const weeksInterval = Math.floor(daysInterval / 7);

	const data: Heatmap = {};

	for (let week = 0; week < weeksInterval; week++) {
		loopDays(week, 7);
	}
	const remainingDays = daysInterval - weeksInterval * 7;
	loopDays(weeksInterval, remainingDays);
	return data;

	function loopDays(w: number, maxDay: number) {
		for (let d = 0; d < maxDay; d++) {
			const date = normalizeDate(new Date((w * 7 + d) * msInDay + start.getTime()));
			let value = 0;
			const sess = sessions.find(
				(s) => normalizeDate(new Date(s.date)).getTime() === date.getTime()
			);
			if (sess) {
				value = sess.time_read;
			}
			if (!data[w]) data[w] = {};
			data[w][d] = {
				date: date.toLocaleDateString('en-GB', { day: '2-digit', month: 'short', year: 'numeric' }),
				time: value,
				level: getActivityLevel(value)
			};
		}
	}
}
function normalizeDate(d: Date) {
	return new Date(Date.UTC(d.getUTCFullYear(), d.getUTCMonth(), d.getUTCDate()));
}

function getActivityLevel(sec: number) {
	if (sec == 0) return 0;
	else if (sec <= 30 * 60) return 1;
	else if (sec <= 60 * 60) return 2;
	else if (sec <= 90 * 60) return 3;
	else if (sec <= 120 * 60) return 4;
	else return 5;
}
