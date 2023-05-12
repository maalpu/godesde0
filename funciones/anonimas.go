package funciones

import "fmt"

// Ej función anónima
func Calculos() {
	num3 := 7
	num4 := 10

	cal := func(num1 int, num2 int) int {
		return num1 + num2 + num3 + num4 // En este caso retornará la suma de todas las variables
	}
	fmt.Println(cal(10, 25))

	cal = func(num1 int, num2 int) int {
		return (num1 * num2 * num3) + num4 // La operatoria es distinta, no así la estructura de la función anónima (cal)
	}
	fmt.Println(cal(5, 2))
}
