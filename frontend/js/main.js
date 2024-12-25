// Cargar lista de productos al cargar la página
document.addEventListener('DOMContentLoaded', async () => {
    try {
        const response = await fetch('/products');
        
        if (!response.ok) {
            throw new Error('Error al cargar los productos');
        }

        const products = await response.json();
        const productsList = document.getElementById('products-list');
        productsList.innerHTML = '';

        products.forEach(product => {
            const productItem = document.createElement('div');
            productItem.className = 'product-item';
            productItem.innerHTML = `
                <h3>${product.name}</h3>
                <p>${product.description}</p>
                <p><strong>Precio: $${product.price}</strong></p>
                <button onclick="addToCart(${product.id})">Agregar al carrito</button>
            `;
            productsList.appendChild(productItem);
        });
    } catch (error) {
        console.error('Error al cargar productos:', error);
        document.getElementById('products-list').textContent = 'Error al cargar los productos';
    }
});

// Función para agregar producto al carrito
function addToCart(productId) {
    let cart = JSON.parse(localStorage.getItem('cart')) || [];
    if (!cart.includes(productId)) {
        cart.push(productId);
        localStorage.setItem('cart', JSON.stringify(cart));
        alert('Producto agregado al carrito');
    } else {
        alert('El producto ya está en el carrito');
    }
}

// Mostrar carrito
function showCart() {
    const cart = JSON.parse(localStorage.getItem('cart')) || [];
    if (cart.length === 0) {
        alert('El carrito está vacío');
    } else {
        alert(`Productos en el carrito: ${cart.join(', ')}`);
    }
}
