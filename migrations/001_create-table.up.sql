CREATE TABLE IF NOT EXISTS split (
    split_id  SERIAL PRIMARY KEY,
    my_percentage REAL NOT NULL,
    her_percentage REAL NOT NULL,
    created_at DATE DEFAULT CURRENT_DATE
);

CREATE TABLE IF NOT EXISTS expense_type (
    type_name VARCHAR(32) PRIMARY KEY,
    active BOOLEAN NOT NULL DEFAULT true,
    created_at DATE DEFAULT CURRENT_DATE
);

CREATE TYPE expense_status_type AS ENUM ('Expensed', 'Unexpensed', 'Covered');

CREATE TABLE IF NOT EXISTS expense (
    expense_id SERIAL PRIMARY KEY,
    expense_type VARCHAR(32) NOT NULL REFERENCES expense_type(type_name),
    amount REAL NOT NULL,
    expense_status expense_status_type DEFAULT 'Unexpensed',
    created_at DATE DEFAULT CURRENT_DATE
);
