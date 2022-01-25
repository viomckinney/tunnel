CREATE TABLE sessions (
    id      UUID NOT NULL PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    token   TEXT NOT NULL UNIQUE
);
