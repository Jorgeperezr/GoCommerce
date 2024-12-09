package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LeerEntrada(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	entrada, _ := reader.ReadString('\n')
	return strings.TrimSpace(entrada)
}
