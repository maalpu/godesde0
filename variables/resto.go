package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Desarrollador string = "Mario"
var Estado bool
var Sueldo float64
var Fecha time.Time

func RestoVariables() {
	Estado = true
	Sueldo = 2580.5
	Fecha = time.Now()
	fmt.Println("Desarrollador = ", Desarrollador)
	fmt.Println("Estado        = ", Estado)
	fmt.Println("Sueldo        = ", Sueldo)
	fmt.Println("Fecha         = ", Fecha)
}

func ConviertoATexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto
}
