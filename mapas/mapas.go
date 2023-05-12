package mapas

import (
	"fmt"
	"strings"
)

func MostrarMapas() {
	pcias := make(map[string]string)

	// fmt.Println(paises)

	pcias["Mendoza"] = "Godoy Cruz"
	pcias["Córdoba"] = "Carlos Paz"
	pcias["San Luis"] = "Villa Mercedes"
	pcias["Tucumán"] = "Tafí del Valle"
	fmt.Println(pcias)

	for provincia, localidad := range pcias {
		fmt.Printf("%s - %s\n", provincia, localidad)
	}
	fmt.Println(strings.Repeat("-", 45))

	delete(pcias, "Mendoza")
	pcias["La Pampa"] = "Santa Rosa"

	for provincia, localidad := range pcias {
		fmt.Printf("%s - %s\n", provincia, localidad)
	}
	fmt.Println(strings.Repeat("-", 45))

	paises := map[string][]string{
		"México":    {"Ciudad de México", "Guadalajara", "Puerto Vallarta"},
		"Argentina": {"Ciudad Autónoma de Buenos Aires", "Buenos Aires", "Mendoza", "Córdoba", "San Juan"},
		"Chile":     {"Santiago", "Valparaiso", "La Serena"},
		"Brasil":    {"Brasilia", "Río de Janeiro", "San Pablo", "Curitiva"},
	}

	fmt.Println(paises)
	fmt.Println(paises["Argentina"])

	campeonato := map[string][]int{
		"Barcelona":   {39, 15, 12, 3, 0, 18, 10},
		"Real Madrid": {38, 15, 12, 2, 1, 16, 11},
		"Chivas":      {37, 15, 12, 1, 2, 18, 13},
		"Boca Junior": {30, 15, 10, 0, 3, 15, 14},
	}

	fmt.Println(campeonato)
	fmt.Println(campeonato["Chivas"])

	// Iteramos el map campeonato
	fmt.Println("Ps  PJ  PG  PE  PP  GF  GC  DI")
	for equipo, p := range campeonato {
		//fmt.Printf("Equipo: %s, tiene %d puntos\n", equipo, puntaje)
		fmt.Printf("%d  %d  %d   %d   %d  %d  %d  %d  | %s \n", p[0], p[1], p[2], p[3], p[4], p[5], p[6], p[5]-p[6], equipo)
	}

	// Iteramos el map paises
	for pais, p := range paises {
		fmt.Printf("%s - %s\n", pais, p[0])
		for i := 1; i < len(p); i++ {
			fmt.Printf("         %s\n", p[i])
		}
	}
	// El comando delete elimina un elemento por la clave
	// Formato: delete(map, clave)
	// Si la clave no es correcta no dará error y por consiguiente no borrará ningún elemento
	delete(paises, "Argento")
	delete(paises, "Brasil")

	fmt.Println(strings.Repeat("-", 45))

	for pais, p := range paises {
		fmt.Printf("%s - %s\n", pais, p[0])
		for i := 1; i < len(p); i++ {
			fmt.Printf("         %s\n", p[i])
		}
	}

	// En los mapas se puede consultar si existe un elemento por su clave
	loca, existe := pcias["San Juan"]
	fmt.Println("Localidad", loca, "  Existe: ", existe)

	loca, existe = pcias["Tucumán"]
	fmt.Printf("Localidad %s - Existe: %t", loca, existe)
}
