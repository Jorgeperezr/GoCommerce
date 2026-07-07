# GoCommerce

**Asignatura:** Programación Orientada a Objetos  
**Estudiante:** Jorge Pérez Rodríguez  
**Carrera:** Sistemas de Información  

---

## 📚 **Descripción del Proyecto**

**GoCommerce** es un sistema de gestión de inventarios y ventas desarrollado en el lenguaje de programación **Go (Golang)**. Este proyecto está diseñado siguiendo los principios de **Programación Orientada a Objetos (POO)**, aplicando encapsulamiento, modularidad, y buenas prácticas en el desarrollo de software.

El sistema permite la gestión eficiente de productos, usuarios y transacciones, garantizando la seguridad mediante autenticación JWT, y una interfaz RESTful clara para interactuar con los servicios.

---

## ⚙️ **Tecnologías Utilizadas**

- **Lenguaje:** Go (Golang)  
- **Framework Web:** Gorilla Mux  
- **Base de Datos:** SQLite  
- **Autenticación:** JWT (JSON Web Tokens)  
- **ORM:** GORM  
- **Herramientas de Seguridad:** Middleware personalizado para validación de tokens y control de acceso  

---

## 🔑 **Funcionalidades Actuales**

### **1. Autenticación y Autorización**
- Inicio de sesión seguro con validación de credenciales.
- Registro de nuevos usuarios con cifrado de contraseñas.
- Generación y renovación de tokens JWT.
- Cierre de sesión seguro.
- Perfil de usuario autenticado.

### **2. Gestión de Productos**
- Listado de todos los productos disponibles.
- Búsqueda de productos por ID.
- Creación de nuevos productos (requiere permisos de administrador).
- Actualización de productos existentes.
- Eliminación de productos.

### **3. Middleware Personalizado**
- Validación de token JWT para proteger rutas privadas.
- Control de acceso basado en roles (`admin`, `user`).
- Registro de solicitudes HTTP para auditoría.

### **4. Gestión de Errores**
- Respuestas claras y estructuradas en formato JSON.
- Manejo de errores personalizados en cada ruta.

---

## 📂 **Estructura del Proyecto**

```plaintext
GoCommerce/
│
├── backend/
│   ├── controllers/   # Controladores para manejar la lógica de negocio
│   ├── models/        # Definiciones de los modelos ORM
│   ├── routers/       # Rutas y endpoints
│   ├── middleware/    # Middleware personalizado
│   ├── utils/         # Funciones auxiliares
│   ├── database/      # Configuración de la base de datos
│   ├── migrations/    # Datos iniciales (usuarios por defecto)
│   ├── tests/         # Pruebas de integración de la API
│   ├── .env.example   # Plantilla de variables de entorno
│   ├── main.go        # Archivo principal para iniciar el servidor
│
├── frontend/
│   ├── auth/          # Página de inicio de sesión
│   ├── css/           # Estilos CSS
│   ├── js/            # Lógica JavaScript
│   ├── index.html     # Página principal
│   ├── admin.html     # Panel de administración
│
├── .devcontainer/     # Configuración para GitHub Codespaces
└── README.md          # Documentación del proyecto
```

---

## 🛠️ **Pasos para Clonar y Ejecutar el Proyecto**

### **1. Clonar el Repositorio**

Abre tu terminal y ejecuta el siguiente comando:

```sh
git clone https://github.com/tu-usuario/GoCommerce.git
cd GoCommerce/backend
```

---

### **2. Configurar las Variables de Entorno**

Copia la plantilla y genera una clave secreta para JWT (el servidor no arranca sin `SECRET_KEY`):

```sh
cp .env.example .env
# Genera una clave y pégala en SECRET_KEY dentro de .env:
openssl rand -hex 32
```

La base de datos SQLite (`gocommerce.db`) se crea automáticamente al arrancar, junto con dos usuarios por defecto:

| Usuario     | Contraseña     | Rol   |
|-------------|----------------|-------|
| `admin`     | `Admin1234`    | admin |
| `empleado1` | `Empleado1234` | user  |

---

### **3. Instalar Dependencias**

Asegúrate de tener Go instalado, luego ejecuta:

```sh
go mod tidy
```

---

### **4. Compilar, Probar y Ejecutar el Proyecto**

```sh
go build ./...   # compila todos los paquetes
go test ./...    # ejecuta las pruebas de integración
go run .         # inicia el servidor (desde backend/)
```

---

### **5. Acceder al Servidor**

Abre tu navegador y visita:

```
http://localhost:8080
```

---

## ✅ **Pruebas de las Rutas Principales**

### **Autenticación:**
- `POST /auth/login`
- `POST /auth/register`
- `POST /auth/refresh-token`
- `POST /auth/logout`
- `GET /auth/profile`

### **Gestión de Productos:**
- `GET /products` (pública)
- `GET /products/{id}` (pública)
- `POST /products` (admin)
- `PUT /products/{id}` (admin)
- `DELETE /products/{id}` (admin)

### **Pedidos:**
- `POST /orders` — crea un pedido; el servidor valida el stock, lo descuenta y calcula el total (usuario autenticado)
- `GET /orders/{id}` (usuario autenticado)
- `GET /orders` (admin)
- `PUT /orders/{id}` — cambia el estado: `pending`, `completed`, `cancelled` (admin)
- `DELETE /orders/{id}` (admin)

### **Administración:**
- `GET /admin/dashboard` — totales de usuarios y productos (admin)
- `GET /admin/transactions` — todas las transacciones (admin)
- `GET /admin/transactions/{user_id}` — historial por usuario (admin)
- `PUT /admin/users/{id}` — actualizar usuario (admin)
- `DELETE /admin/users/{id}` — eliminar usuario (admin)

---

## 📑 **Notas Importantes**

- Asegúrate de que el archivo `.env` esté configurado correctamente.
- Verifica que la base de datos SQLite esté en el directorio raíz del backend.
- Para acceder a las rutas protegidas, incluye un token JWT válido en el encabezado `Authorization`.

---

## 🤝 **Contribuciones**

Las contribuciones son bienvenidas. Para realizar mejoras o sugerencias, abre un **Issue** o un **Pull Request** en el repositorio.

---

## 📄 **Licencia**

Este proyecto está bajo la licencia **MIT**.
