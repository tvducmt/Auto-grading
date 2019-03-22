

CREATE TABLE users (
    id          SERIAL PRIMARY KEY,
    username   TEXT NOT NULL UNIQUE,
    email       TEXT NOT NULL UNIQUE,
    passphrase  TEXT NOT NULL,
    fullname    TEXT,
    role       TEXT DEFAULT false NOT NULL,
    is_delete BOOLEAN DEFAULT false NOT NULL
);
