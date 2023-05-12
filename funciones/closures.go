package funciones

import "fmt"

func tabla(valor int) func() int {
	num := valor
	secuencia := 0
	return func() int {
		secuencia++
		return num * secuencia
	}
}

func LlamarClosure() {
	tablaDel := 2
	miTabla := tabla(tablaDel)
	for i := 1; i <= 10; i++ {
		fmt.Println(miTabla())
	}
}
