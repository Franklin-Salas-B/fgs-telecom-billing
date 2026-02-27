#!/bin/bash

# Configuración de MySQL
DB_USER="root"
DB_PASS="P1ch1nch4"  # tu contraseña de root
DB_NAME="ecommerce_db"

echo "✅ Iniciando prueba automatizada de ecommerce..."

# 1️⃣ Insertar usuario de prueba
mysql -u $DB_USER -p$DB_PASS $DB_NAME <<EOF
INSERT INTO usuarios (nombre, email) VALUES ('Usuario Test', 'test@example.com') 
ON DUPLICATE KEY UPDATE nombre=nombre;
EOF
echo "Usuario de prueba insertado."

# 2️⃣ Insertar productos de prueba
mysql -u $DB_USER -p$DB_PASS $DB_NAME <<EOF
INSERT INTO productos (nombre, descripcion, precio, creado_en) VALUES
('Producto A', 'Descripción A', 10.00, '2026-02-27'),
('Producto B', 'Descripción B', 15.50, '2026-02-27'),
('Producto C', 'Descripción C', 7.25, '2026-02-27')
ON DUPLICATE KEY UPDATE nombre=nombre;
EOF
echo "Productos de prueba insertados."

# 3️⃣ Limpiar carrito antes de prueba
mysql -u $DB_USER -p$DB_PASS $DB_NAME <<EOF
DELETE FROM carrito;
EOF
echo "Carrito limpiado."

# 4️⃣ Llenar carrito con varios productos
mysql -u $DB_USER -p$DB_PASS $DB_NAME <<EOF
INSERT INTO carrito (usuario_id, producto_id, cantidad, creado_en) VALUES
(1, 1, 2, '2026-02-27'),
(1, 2, 1, '2026-02-27'),
(1, 3, 3, '2026-02-27');
EOF
echo "Carrito llenado con 3 productos."

# 5️⃣ Ejecutar checkout mediante Go API
echo "Realizando checkout..."
curl -s -X POST -H "Content-Type: application/json" \
-d '{"usuario_id":1}' \
http://localhost:8081/checkout | jq .

# 6️⃣ Mostrar carrito después de checkout
echo "Carrito después de checkout:"
curl -s http://localhost:8081/carrito | jq .

# 7️⃣ Mostrar órdenes
echo "Órdenes creadas:"
curl -s http://localhost:8081/ordenes | jq .

# 8️⃣ Mostrar detalle de órdenes
echo "Detalle de órdenes (MySQL):"
mysql -u $DB_USER -p$DB_PASS $DB_NAME <<EOF
SELECT d.id, d.orden_id, o.usuario_id, d.producto_id, d.cantidad, d.precio
FROM detalle_orden d
JOIN ordenes o ON d.orden_id = o.id;
EOF

echo "✅ Prueba completada."
