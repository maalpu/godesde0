package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func IngreseValor() string {
	var tabla string
	for {
		valor := bufio.NewScanner(os.Stdin)
		fmt.Print("Ingrese numero: ")
		if valor.Scan() {
			numero, err := strconv.Atoi(valor.Text())
			if err == nil {
				for i := 1; i <= 10; i++ {
					tabla += fmt.Sprintf("%d x %d = %d\n", numero, i, numero*i)
				}
				return tabla
			}
		}
	}
}
