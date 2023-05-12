package arreglos_slices

import "fmt"

// A diferencia de un array, en un slice no se debe definir la cantidad de elementos
// a un slice se le pueden agregar y quitar elementos
var tablaS []int = []int{2, 5, 6}

var arreglo [10]int = [10]int{6, 98, 21, 674, 18, 36, 78, 9}

func MuestroSlice() {
	fmt.Println(tablaS)

	porcion := arreglo[3:] // Creamos un slice desde un array desde la posición 3

	porcion2 := arreglo[:5] // Creamos un slice desde un array hasta la posición anterior a 5

	porcion3 := arreglo[6:7] // Creamos un slice de la posición 6

	fmt.Println("Porcion :", porcion)
	fmt.Println("Porcion2:", porcion2)
	fmt.Println("Porcion3:", porcion3)
}

func Capacidad() {
	// make permite definir la capacidad de un slice
	algo := make([]int, 5, 20) // Definimos un slice de 5 elementos y 20 de Capacidad

	fmt.Printf("len(algo): %d (Longitud)   -   cap(algo): %d (Capacidad) \n", len(algo), cap(algo))

	// Definimos nums como slice sin longitud, ni capacidad
	nums := make([]int, 0, 0)
	fmt.Printf("len(nums): %d (Longitud)   -   cap(nums): %d (Capacidad)   -   Definimos nums como slice sin longitud, ni capacidad\n", len(nums), cap(nums))

	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	// Go researva memoria de capacidad por múltiplos de 2
	fmt.Printf("len(nums): %d (Longitud)   -   cap(nums): %d (Capacidad) \n", len(nums), cap(nums))
}
