package main

import (
	"fmt"

	"github.com/maalpu/godesde0/variables"
)

func main() {
	// variables.MuestroEnteros()
	// variables.RestoVariables()
	Estado, texto := variables.ConviertoATexto(13035)
	fmt.Println(texto, Estado)
}
