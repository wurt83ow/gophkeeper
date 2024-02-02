CREATE TABLE IF NOT EXISTS UserCredentials (
    id TEXT PRIMARY KEY,
    user_id INTEGER,
    login TEXT NOT NULL,
    password TEXT NOT NULL,
    meta_info TEXT,
    deleted BOOLEAN DEFAULT FALSE,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(user_id) REFERENCES Users(id)
);