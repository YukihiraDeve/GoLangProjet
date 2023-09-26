package main

import "fmt"

func main() {
	type Ville struct {
		name       string
		codePostal int
		superficie int
		population int
	}

	// Création de la map avec des clés de type string et des valeurs de type int
	var Value = make(map[string]int)

	// Saisie des valeurs dans des variables temporaires, puis ajout dans la map
	var temp int
	fmt.Println("Enter value  : Code postal")
	fmt.Scanf("%d", &temp)
	Value["CodePostal"] = temp

	fmt.Println("Enter value  : superficie")
	fmt.Scanf("%d", &temp)
	Value["Superficie"] = temp

	fmt.Println("Enter value  : population")
	fmt.Scanf("%d", &temp)
	Value["Population"] = temp

	// Création d'une instance de Ville à partir des valeurs dans la map
	ville := Ville{
		codePostal: Value["CodePostal"],
		superficie: Value["Superficie"],
		population: Value["Population"],
	}

	// Affichage des valeurs de la structure Ville
	fmt.Println("Code Postal:", ville.codePostal)
	fmt.Println("Superficie:", ville.superficie)
	fmt.Println("Population:", ville.population)
}
