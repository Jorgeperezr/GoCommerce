document.getElementById('loginForm').addEventListener('submit', async (e) => {
    e.preventDefault();

    const usuario = document.getElementById('usuario').value;
    const contrasena = document.getElementById('contrasena').value;
    const error = document.getElementById('error');

    try {
        const response = await fetch('/api/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ usuario, contrasena }),
        });

        if (!response.ok) {
            throw new Error('Credenciales incorrectas');
        }

        window.location.href = '/dashboard.html';
    } catch (err) {
        error.textContent = err.message;
    }
});
