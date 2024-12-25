// Verificar token al cargar la página
document.addEventListener('DOMContentLoaded', async () => {
    const token = localStorage.getItem('token');
    if (!token) {
        window.location.href = 'login.html';
        return;
    }

    try {
        const response = await fetch('/auth/profile', {
            headers: { 'Authorization': `Bearer ${token}` }
        });
        
        if (!response.ok) {
            throw new Error('Token inválido o expirado');
        }

        const userData = await response.json();
        document.querySelector('header h1').textContent = `Bienvenido, ${userData.username}`;
    } catch (error) {
        console.error('Error al validar el token:', error);
        localStorage.removeItem('token');
        window.location.href = 'login.html';
    }
});

// Cerrar sesión
document.getElementById('logoutBtn').addEventListener('click', () => {
    localStorage.removeItem('token');
    window.location.href = 'login.html';
});

// Cargar datos en el dashboard
async function loadDashboard() {
    try {
        const response = await fetch('/dashboard', {
            headers: { 'Authorization': `Bearer ${localStorage.getItem('token')}` }
        });
        
        if (!response.ok) {
            throw new Error('Error al cargar el dashboard');
        }

        const data = await response.json();
        document.getElementById('dashboard-content').innerHTML = `Usuarios: ${data.total_users} | Productos: ${data.total_products}`;
    } catch (error) {
        console.error('Error al cargar el dashboard:', error);
        document.getElementById('dashboard-content').textContent = 'Error al cargar el dashboard';
    }
}

// Llamar funciones iniciales
loadDashboard();
