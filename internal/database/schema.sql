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
	authors TEXT NOT NULL,
	pages INTEGER NOT NULL,
	description TEXT,
    publisher TEXT,
    publish_date TEXT,
    language TEXT NOT NULL,
    cover_url TEXT 
);

CREATE TABLE IF NOT EXISTS reading (
    user_id TEXT,
    book_id TEXT,
    pages_read INTEGER DEFAULT 0,
    time_read INTEGER DEFAULT 0,
    first_page INTEGER DEFAULT 1,
    current_page INTEGER DEFAULT 0,
    session_created_at TEXT,
    session_updated_at TEXT,
    PRIMARY KEY (user_id, book_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (book_id) REFERENCES books(isbn) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS groups (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS members (
    user_id TEXT,
    group_id TEXT,
    joined_at TEXT NOT NULL,
    PRIMARY KEY (user_id, group_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS invitations (
    email_to TEXT NOT NULL,
    group_id TEXT NOT NULL,
    invited_by TEXT NOT NULL,
    status TEXT DEFAULT 'pending',
    sent_at TEXT NOT NULL,
    expires_at TEXT NOT NULL,
    PRIMARY KEY (email_to, group_id),
    FOREIGN KEY (email_to) REFERENCES users(email) ON DELETE CASCADE,
    FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE,
    FOREIGN KEY (invited_by) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS stats (
    user_id TEXT PRIMARY KEY,
    total_pages INTEGER DEFAULT 0,
    total_time INTEGER DEFAULT 0,
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS sessions {
    user_id TEXT NOT NULL,
    date TEXT NOT NULL,
    time_read INTEGER DEFAULT 0,
    PRIMARY KEY (user_id, date),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
};
