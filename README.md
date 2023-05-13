# Inventario de productos electrodomesticos
Aplicacion construida con arquitectura hexagonal para poder separar la logica y el dominio del negocio de la infraestructura a traves de puertos o interfaces
## Modelos de negocio:
- Producto:
    - Nombre
    - Descripcion
    - Codigo de barras (identificador de lote)
    - Cantidad disponible
    - Precio
    - Proveedor
- Proveedor:
    - Nombre de compañia
    - Region

## Logica de negocio:
1. El dueño de la empresa va poder registrar una entrega de un lote de equis cantidad de electrodomesticos por parte de un proveedor con todos los campos que se especifican en el modelo de producto
2. El campo codigo de barras debe ser unico, y de haber una entrega con el mismo codigo de barras saltaria un error
3. El precio del producto debe ser distinto de 0 y un numero positivo

ESTA APLICACION SOLO ES PARA FINES EDUCATIVOS Y NO DEBE TOMARSE COMO UNA APLICACION DE INVENTARIADO REAL
