document.getElementById('loginForm').addEventListener('submit', async (e) => {
    e.preventDefault();
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;
    const errorMessage = document.getElementById('errorMessage');
    errorMessage.textContent = '';

    try {
        const response = await fetch('/auth/login', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ username, password })
        });

        if (response.ok) {
            const data = await response.json();
            localStorage.setItem('token', data.token);
            window.location.href = 'admin.html';
        } else if (response.status === 401) {
            errorMessage.textContent = 'Credenciales incorrectas. Inténtalo de nuevo.';
        } else if (response.status === 400) {
            errorMessage.textContent = 'Solicitud incorrecta. Verifica tus datos.';
        } else {
            errorMessage.textContent = 'Error en el servidor. Inténtalo más tarde.';
        }
    } catch (error) {
        console.error('Error en autenticación:', error);
        errorMessage.textContent = 'Error inesperado. Inténtalo de nuevo.';
    }
});

// Validar token al cargar la página
document.addEventListener('DOMContentLoaded', () => {
    const token = localStorage.getItem('token');
    if (token) {
        fetch('/auth/profile', {
            headers: { 'Authorization': `Bearer ${token}` }
        })
        .then(response => {
            if (!response.ok) {
                throw new Error('Token inválido');
            }
            return response.json();
        })
        .then(data => {
            console.log('Usuario autenticado:', data);
        })
        .catch(error => {
            console.error('Error al validar el token:', error);
            localStorage.removeItem('token');
            window.location.href = 'login.html';
        });
    }
});
