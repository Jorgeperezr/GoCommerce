// Package utils proporciona funciones auxiliares.
package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// LeerEntrada lee una entrada desde la consola.
func LeerEntrada(mensaje string) string {
	fmt.Print(mensaje)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

// LeerEntero lee un número entero desde la consola.
func LeerEntero(mensaje string) (int, error) {
	entrada := LeerEntrada(mensaje)
	return strconv.Atoi(entrada)
}
