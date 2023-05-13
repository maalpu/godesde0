package middleware

import "fmt"

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func dividir(a, b int) int {
	return a / b
}

func MiMiddleware() {
	fmt.Println("Inicio Middleware")

	result := operacionesMidd(sumar)(2, 3)
	fmt.Printf("Sumar %d+%d........: %d\n", 2, 3, result)

	result = operacionesMidd(restar)(10, 6)
	fmt.Printf("Restar %d-%d......: %d\n", 10, 6, result)

	result = operacionesMidd(multiplicar)(2, 4)
	fmt.Printf("Multiplicar %d*%d..: %d\n", 2, 4, result)

	result = operacionesMidd(dividir)(21, 3)
	fmt.Printf("Dividir %d/%d.....: %d\n", 21, 3, result)
}

func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("                       Inicio de Operación") // Acá podríamos validar y por consiguiente abortar o seguir
		return f(a, b)
	}
}
