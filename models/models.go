package models

// Producto representa un producto
type Producto struct {
	ID        int
	Nombre    string
	Precio    float64
	Stock     int
	Categoria string
}

// Historial representa una transacción
type Historial struct {
	ID       int
	Producto string
	Usuario  string
	Tipo     string
	Cantidad int
	Fecha    string
}
