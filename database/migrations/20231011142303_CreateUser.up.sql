CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    document VARCHAR UNIQUE,
    email VARCHAR UNIQUE,
    password VARCHAR,
    role VARCHAR
);

CREATE INDEX IF NOT EXISTS "users_id" ON "users" ("id")