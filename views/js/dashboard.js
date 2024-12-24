async function cargarProductos() {
    const res = await fetch('/productos');
    const data = await res.json();
    let tabla = document.getElementById('tabla-productos');
    tabla.innerHTML = '';

    data.forEach(p => {
        tabla.innerHTML += `
      <tr>
        <td>${p.ID}</td>
        <td>${p.Nombre}</td>
        <td>${p.Precio}</td>
        <td>${p.Stock}</td>
        <td>${p.Categoria}</td>
        <td>
          <button onclick="eliminarProducto(${p.ID})">Eliminar</button>
        </td>
      </tr>`;
    });
}

async function eliminarProducto(id) {
    const res = await fetch(`/productos/eliminar?id=${id}`, { method: 'DELETE' });
    if (res.ok) {
        alert('Producto eliminado');
        cargarProductos();
    }
}
