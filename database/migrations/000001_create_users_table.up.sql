CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    dob DATE,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);
