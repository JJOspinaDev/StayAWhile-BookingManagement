-- Crear tabla clientes
CREATE TABLE clientes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100),
    email VARCHAR(100) UNIQUE,
    telefono VARCHAR(50)
);

-- Crear tabla habitaciones
CREATE TABLE habitaciones (
    id INT AUTO_INCREMENT PRIMARY KEY,
    tipo VARCHAR(50),
    descripcion TEXT,
    precioPorNoche DECIMAL(10, 2)
);

-- Crear tabla reservas
CREATE TABLE reservas (
    id INT AUTO_INCREMENT PRIMARY KEY,
    clienteId INT,
    habitacionId INT,
    fechaEntrada DATE,
    fechaSalida DATE,
    numeroNoches INT,
    costoTotal DECIMAL(10, 2),
    estado VARCHAR(50),
    desayunoIncluido BOOLEAN,
    camaExtra BOOLEAN,
    transporteAeropuerto BOOLEAN,
    notas TEXT,
    fechaCreacion DATETIME,
    fechaActualizacion DATETIME,
    FOREIGN KEY (clienteId) REFERENCES clientes(id),
    FOREIGN KEY (habitacionId) REFERENCES habitaciones(id)
);

-- Insertar datos de prueba para clientes
INSERT INTO clientes (nombre, email, telefono) VALUES 
('Ana Torres', 'ana.torres@example.com', '1234567890'),
('Carlos Gomez', 'carlos.gomez@example.com', '1234567891'),
('Lucia Hernandez', 'lucia.hernandez@example.com', '1234567892'),
('David Jimenez', 'david.jimenez@example.com', '1234567893'),
('Sofia Lopez', 'sofia.lopez@example.com', '1234567894');


-- Insertar datos de prueba para habitaciones
INSERT INTO habitaciones (tipo, descripcion, precioPorNoche) VALUES 
('Individual', 'Habitacion individual estandar', 50.00),
('Individual', 'Habitacion individual estandar', 50.00),
('Individual', 'Habitacion individual estandar', 50.00),
('Doble', 'Habitacion doble economica', 60.00),
('Doble', 'Habitacion doble estandar', 65.00),
('Doble', 'Habitacion doble estandar', 65.00),
('Suite', 'Suite de lujo con vistas panoramicas', 180.00),
('Suite', 'Suite de lujo con vistas panoramicas', 180.00),
('Suite', 'Suite presidencial con jacuzzi', 200.00);


-- Insertar datos de prueba para reservas
INSERT INTO reservas (clienteId, habitacionId, fechaEntrada, fechaSalida, numeroNoches, costoTotal, estado, desayunoIncluido, camaExtra, transporteAeropuerto, notas, fechaCreacion, fechaActualizacion) VALUES 
(1, 1, '2023-01-01', '2023-01-03', 2, 100.00, 'Confirmada', TRUE, FALSE, FALSE, 'Sin notas', NOW(), NOW()),
(2, 2, '2023-01-05', '2023-01-10', 5, 350.00, 'Confirmada', TRUE, TRUE, TRUE, 'Requiere transporte aeropuerto', NOW(), NOW());
-- Añadir más reservas hasta llegar a 13...
