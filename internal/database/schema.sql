CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    pw_hash TEXT NOT NULL,
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL
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

CREATE TABLE IF NOT EXISTS views (
    user_id TEXT,
    book_id TEXT,
    PRIMARY KEY (user_id, book_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (book_id) REFERENCES books(isbn) ON DELETE CASCADE
);