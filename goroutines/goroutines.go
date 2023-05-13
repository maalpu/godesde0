package goroutines

import (
	"fmt"
	"time"
)

func MiNombreLento(nombre string, canal chan bool) {
	defer fmt.Println()
	for i := 0; i < len(nombre); i++ {
		// Haremos una pausa de un segundo por letra
		fmt.Print(nombre[i : i+1])
		time.Sleep(300 * time.Millisecond)
	}
	canal <- true // Así se le asigna un valor al (chan) canal
}
