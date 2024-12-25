package utils

// Roles de usuario
const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)

// Estados de pedido
const (
	OrderStatusPending   = "pending"
	OrderStatusCompleted = "completed"
	OrderStatusCancelled = "cancelled"
)

// Mensajes de error comunes
const (
	ErrUnauthorized   = "No autorizado"
	ErrForbidden      = "Acceso denegado"
	ErrNotFound       = "Recurso no encontrado"
	ErrInvalidInput   = "Entrada inválida"
	ErrInternalServer = "Error interno del servidor"
)

// Configuraciones generales
const (
	TokenExpirationHours = 24
	DefaultPageSize      = 20
)
