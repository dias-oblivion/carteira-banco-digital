CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    document VARCHAR(20) UNIQUE,
    email VARCHAR(100) UNIQUE,
    password VARCHAR(100),
    role VARCHAR(10),
    balance NUMERIC,
    created_at TIMESTAMP
);
CREATE INDEX IF NOT EXISTS "users_id" ON "users" ("id");

CREATE TABLE IF NOT EXISTS transfers_history (
    id SERIAL PRIMARY KEY,
    fk_payer_id INT REFERENCES users(id),
    fk_payee_id INT REFERENCES users(id),
    description VARCHAR(100),
    value NUMERIC
);
CREATE INDEX IF NOT EXISTS "transfers_history_id" ON "transfers_history" ("id");