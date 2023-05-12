package arreglos_slices

import "fmt"

// Un array o arreglo se debe definir con la cantidad de elementos que tendrá y no puede ser ajustado
// una vez declarado un array se reservan los elementos en memoria declarados
var tabla [10]int // Creo un array de 10 elementos del [0-9]

var tabla2 [12]int = [12]int{2, 16, 914} // Creamos un array y le asignamos valores

var matriz [20][30]int

func MuestraArreglos() {
	tabla[7] = 33
	tabla[2] = 54
	fmt.Println(tabla)

	tabla2[10] = 4562
	fmt.Println(tabla2)

	tabla3 := [10]string{"Juan", "Luisa"} // Creamos un array y le asignamos valores
	tabla3[6] = "Marisa"
	fmt.Println(tabla3)

	// Iterar un array
	for i := 0; i < len(tabla3); i++ {
		fmt.Println("i =", i, " - ", tabla3[i])
	}

	matriz[7][24] = 15
	fmt.Println(matriz)
}
