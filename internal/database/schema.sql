CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS books (
	isbn TEXT PRIMARY KEY,
	title TEXT NOT NULL UNIQUE,
	author TEXT NOT NULL,
	pages INTEGER NOT NULL,
	description TEXT,
    publisher TEXT,
    publish_date TEXT,
    language INTEGER NOT NULL CHECK (language BETWEEN 0 AND 2)
);