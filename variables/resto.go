package variables

import (
	"fmt"
	"time"
)

var Desarrollador string = "Mario A. Puebla"
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