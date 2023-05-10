package ejercicios

import "strconv"

func ComprobarValor(cadena string) (string, int) {
	valEntero, err := strconv.Atoi(cadena)
	if err != nil {
		return "Error:  " + err.Error(), 0
	}
	if valEntero > 100 {
		return "Es mayoy a 100", valEntero
	} else {
		return "Es menor o igual a 100", valEntero
	}
}
