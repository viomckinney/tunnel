CREATE TABLE users (
    id             UUID  NOT NULL PRIMARY KEY,
    username       TEXT  NOT NULL UNIQUE,
    username_lower TEXT  NOT NULL UNIQUE,
    password_hash  BYTEA NOT NULL
);

CREATE TABLE api_keys (
    id       UUID NOT NULL PRIMARY KEY,
    user_id  UUID NOT NULL REFERENCES users(id),
    token    TEXT NOT NULL UNIQUE
);

CREATE TABLE invite_codes (
    id   UUID NOT NULL PRIMARY KEY,
    code TEXT NOT NULL UNIQUE
);
