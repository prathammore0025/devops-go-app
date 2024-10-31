-- init.sql
CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    employee_id INT NOT NULL
);

-- You can add more tables or initial data here if needed.

