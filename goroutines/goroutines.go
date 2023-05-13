package goroutines

import (
	"fmt"
	"time"
)

func MiNombreLento(nombre string) {
	defer fmt.Println()
	for i := 0; i < len(nombre); i++ {
		// Haremos una pausa de un segundo por letra
		fmt.Print(nombre[i : i+1])
		time.Sleep(300 * time.Millisecond)
	}
}
