CREATE TABLE users (
                       id SERIAL NOT NULL, UNIQUE,
                       name VARCHAR(255) NOT NULL,
                       last_name VARCHAR(255) NOT NULL,
                       second_name VARCHAR(255) NOT NULL,
                       password_hash VARCHAR(255) NOT NULL,
                       role VARCHAR(255) NOT NULL
);
CREATE TABLE products (
                          id SERIAL NOT NULL, UNIQUE,
                          name VARCHAR(255) NOT NULL,
                          weight INT NOT NULL,
                          cost DECIMAL NOT NULL
);
