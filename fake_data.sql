-- Insertar proveedores
INSERT INTO providers (name, region) VALUES
('Proveedor 1', 'Región 1'),
('Proveedor 2', 'Región 2'),
('Proveedor 3', 'Región 1'),
('Proveedor 4', 'Región 3'),
('Proveedor 5', 'Región 2');

-- Insertar productos
INSERT INTO products (name, description, stock, price, provider_id) VALUES
('Producto 1', 'Descripción del producto 1', 10, 9.99, 1),
('Producto 2', 'Descripción del producto 2', 5, 19.99, 1),
('Producto 3', 'Descripción del producto 3', 2, 4.99, 2),
('Producto 4', 'Descripción del producto 4', 8, 14.99, 3),
('Producto 5', 'Descripción del producto 5', 3, 7.99, 2),
('Producto 6', 'Descripción del producto 6', 0, 29.99, 4),
('Producto 7', 'Descripción del producto 7', 12, 12.99, 5),
('Producto 8', 'Descripción del producto 8', 6, 9.99, 4),
('Producto 9', 'Descripción del producto 9', 4, 6.99, 3),
('Producto 10', 'Descripción del producto 10', 7, 24.99, 5);