package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func IngreseValor() {
	for {
		valor := bufio.NewScanner(os.Stdin)
		fmt.Println("Ingrese numero1: ")
		if valor.Scan() {
			numero1, err := strconv.Atoi(valor.Text())
			if err == nil {
				for i := 1; i <= 10; i++ {
					fmt.Printf("%d x %d = %d\n", numero1, i, numero1*i)
				}
				break
			}
		}
	}
}
