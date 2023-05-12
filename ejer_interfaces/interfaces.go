package ejer_interfaces

import (
	"fmt"

	"github.com/maalpu/godesde0/interfaces"
)

func HumanoRespirando(hu interfaces.Humano) {
	hu.Respirar()
	que := ""
	if hu.Sexo() == "Hombre" {
		que = "un"
	} else {
		que = "una"
	}
	fmt.Printf("Soy %s %s, y estoy respirando \n", que, hu.Sexo())
}
