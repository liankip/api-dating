CREATE TABLE premium_packages (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    price DECIMAL(10, 2) NOT NULL
);

ALTER TABLE users ADD COLUMN premium_package_id INT REFERENCES premium_packages(id);