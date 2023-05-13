package main

import (
	"fmt"

	"github.com/maalpu/godesde0/goroutines"
)

func main() {
	// variables.MuestroEnteros()
	// variables.RestoVariables()

	// Estado, texto := variables.ConviertoATexto(13035)
	// fmt.Println(texto, Estado)

	// os := runtime.GOOS
	// if os == "Linux" || os == "OS X" {
	// 	fmt.Println("Esto es Windows")
	// } else {
	// 	fmt.Println("Esto no es Windows, es", os)
	// }
	// fmt.Println("")

	// o := runtime.GOOS
	// fmt.Println()
	// switch o {
	// case "Linux":
	// 	fmt.Println("Esto es Linux")
	// case "OS X":
	// 	fmt.Println("OS X")
	// case "darwin":
	// 	fmt.Println("Darwin")
	// case "Windows":
	// 	fmt.Println("Windows")
	// default:
	// 	fmt.Println("Error")
	// }

	// Esto equvale a lo anterior (Aunque en Linux no funciona igual)
	// if os := runtime.GOOS; os == "Linux." || os == "OS X." {
	// 	fmt.Println("Esto no es Windows")
	// } else {
	// 	fmt.Println("Esto es Windows")
	// }

	// cadena, valor := ejercicios.ComprobarValor("78")
	// fmt.Printf("%d %s\n", valor, cadena)
	// cadena, valor = ejercicios.ComprobarValor("1950")
	// fmt.Printf("%d %s\n", valor, cadena)
	// cadena, valor = ejercicios.ComprobarValor("asdfads")
	// fmt.Printf("%d %s\n", valor, cadena)

	// teclado.IngresoNumeros()

	// iteraciones.Iterar()

	// fmt.Println(ejercicios.IngreseValor())

	// archivos.GrabarTabla()

	// archivos.SumarTabla()

	// archivos.LeoArchivo()
	// archivos.LeoArchivoIOutil()

	// funciones.Calculos()

	// funciones.LlamarClosure()

	// funciones.Exponencia(2)

	// arreglos_slices.MuestraArreglos()

	// arreglos_slices.MuestroSlice()

	// arreglos_slices.Capacidad()

	// mapas.MostrarMapas()

	// users.Alta(1, "Mario Alberto Puebla")

	// Pedro := new(modelos.Hombre)
	// ejer_interfaces.HumanoRespirando(Pedro)

	// Luisa := new(modelos.Mujer)
	// ejer_interfaces.HumanoRespirando(Luisa)

	// defer_panic.VemosDefer()
	// defer_panic.EjemploPanic()

	canal1 := make(chan bool)
	defer func() { <-canal1 }()

	go goroutines.MiNombreLento("Mario Alberto Puebla", canal1)
	fmt.Println("Estamos ejecutando el programa... ")

}
