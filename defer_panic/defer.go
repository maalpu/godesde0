package defer_panic

import (
	"fmt"
	"log"
)

func VemosDefer() {
	// Defer, es una instrucción que nos da Go, que nos permite configurar cual será la última instrucción
	// en ejecutarse cuando salga de una función. Sólo puede ejecutar una instrucción
	// Ejemplo:
	// 		defer archivo.Close() // Nos aseguramos de cerrar el archivo al salir de la función
	fmt.Println("Este es un primer mensaje")
	defer fmt.Println("Este es el mensaje final") // Si queremos ejecutar varias instrucciones, deberíamos utilizar una func

	fmt.Println("Este es el segundo mensaje")

	/* Esto será la salida por pantalla
	Este es un primer mensaje
	Este es el segundo mensaje
	Este es el mensaje final
	*/
}

func EjemploPanic() {
	// panic se utiliza para abortar el sistema
	// El recover sólo puede ser invocado por defer
	// recover no puede capturar un panic, sin utilizar defer
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrió un error que generó panic \n%v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Se encontró el valor 1")
	}
}
