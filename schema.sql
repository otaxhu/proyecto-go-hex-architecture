DROP DATABASE IF EXISTS electro_inventory_app;
CREATE DATABASE electro_inventory_app;
USE electro_inventory_app;

CREATE TABLE providers(
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    region VARCHAR(255) NOT NULL
);

CREATE TABLE products(
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(1000),
    stock INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    provider_id INT NOT NULL,
    FOREIGN KEY(provider_id) REFERENCES providers(id)
);
