CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    document VARCHAR UNIQUE(20),
    email VARCHAR UNIQUE(50),
    password VARCHAR(16),
    role VARCHAR(10),
    balance NUMERIC,
    created_at TIMESTAMP
);
CREATE INDEX IF NOT EXISTS "users_id" ON "users" ("id")

CREATE TABLE IF NOT EXISTS transactions_history (
    id SERIAL PRIMARY_KEY,
    fk_payer_id INT REFERENCES users(id),
    fk_payee_id INT REFERENCES users(id),
    description VARCHAR(100),
    value NUMERIC
)
CREATE INDEX IF NOT EXISTS "transactions_history_id" ON "transactions_history" ("id")