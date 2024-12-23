package api

import (
	"fmt"
	"net/http"
)

func IniciarServidor() {
	http.HandleFunc("/productos", FuncProductos)
	http.HandleFunc("/login", FuncLogin)

	fmt.Println("Servidor escuchando en el puerto 8080")
	http.ListenAndServe(":8080", nil)
}
