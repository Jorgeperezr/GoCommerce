async function cargarProductos() {
    const respuesta = await fetch('/api/productos');
    const productos = await respuesta.json();

    const tbody = document.getElementById('productos');
    tbody.innerHTML = ''; // Limpiar tabla antes de renderizar

    productos.forEach(p => {
        tbody.innerHTML += `
            <tr>
                <td>${p.id}</td>
                <td>${p.nombre}</td>
                <td>${p.precio}</td>
                <td>${p.stock}</td>
                <td>${p.categoria_id}</td>
                <td>
                    <button onclick="eliminarProducto(${p.id})">Eliminar</button>
                </td>
            </tr>
        `;
    });
}

async function agregarProducto() {
    const nombre = document.getElementById('nombre').value;
    const precio = document.getElementById('precio').value;
    const stock = document.getElementById('stock').value;
    const categoria_id = document.getElementById('categoria_id').value;

    await fetch('/api/productos/agregar', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ nombre, precio, stock, categoria_id })
    });

    cargarProductos();
}

async function eliminarProducto(id) {
    await fetch(`/api/productos/eliminar/${id}`, { method: 'DELETE' });
    cargarProductos();
}

document.addEventListener('DOMContentLoaded', cargarProductos);
