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

-- Example insert (customer_id is auto-generated)
INSERT INTO customers (name, date_of_birth, city, zipcode, status)
VALUES ('John Doe', '1985-06-15', 'New York', '10001', 1);
