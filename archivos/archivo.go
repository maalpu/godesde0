package archivos

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/maalpu/godesde0/ejercicios"
)

var fileName string = "./archivos/txt/tablas.txt"

func GrabarTabla() {
	archivo, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error al crear el archivo", err.Error())
		return
	}
	fmt.Fprintln(archivo, fileName)
	archivo.Close()
}

func SumarTabla() {
	tabla := ejercicios.IngreseValor()
	if !append(tabla) {
		fmt.Println("Error al concatenar Tabla")
	}
}

func append(tabla string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		archivo, err := os.Create(fileName)
		if err == nil {
			fmt.Fprintln(archivo, tabla)
			archivo.Close()
			return true
		}
	}

	_, err = arch.WriteString(tabla + "\n")
	if err != nil {
		fmt.Println("Error durante el WriteString ", err.Error())
		return false
	}
	arch.Close()
	return true
}

func LeoArchivoIOutil() bool {
	archivo, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error al leer archivo ", err.Error())
		return false
	}
	fmt.Println(string(archivo))
	return true
}

func LeoArchivo() {
	archivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error al leer archivo", err.Error())
	}
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("| " + registro)
	}
	archivo.Close()
}
