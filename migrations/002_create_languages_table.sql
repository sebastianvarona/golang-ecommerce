CREATE TABLE languages (
    id SERIAL PRIMARY KEY,
    code VARCHAR(10) UNIQUE NOT NULL,  -- e.g., 'en', 'es', 'fr'
    name VARCHAR(100) NOT NULL
);
