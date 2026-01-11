-- Create the people table
CREATE TABLE IF NOT EXISTS people (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

INSERT INTO people (name) VALUES ('Bob'), ('Charlie');

-- PostgreSQL version of the customers table
CREATE TABLE customers (
    customer_id SERIAL PRIMARY KEY,   -- auto-incrementing ID
    name VARCHAR(100) NOT NULL,
    date_of_birth DATE NOT NULL,
    city VARCHAR(100) NOT NULL,
    zipcode VARCHAR(10) NOT NULL,
    status SMALLINT NOT NULL          -- tinyint in MySQL → SMALLINT in Postgres
);

CREATE TABLE accounts (
    account_id SERIAL PRIMARY KEY,           -- auto-incrementing account ID
    customer_id INT NOT NULL,                -- foreign key to customers
    opening_date TIMESTAMP NOT NULL,
    account_type VARCHAR(10) NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    status SMALLINT NOT NULL,

    -- Foreign key constraint
    CONSTRAINT fk_customer
        FOREIGN KEY (customer_id)
        REFERENCES customers(customer_id)
        ON DELETE CASCADE
);

-- Example insert (customer_id is auto-generated)
INSERT INTO customers (name, date_of_birth, city, zipcode, status)
VALUES ('John Doe', '1985-06-15', 'New York', '10001', 1);
