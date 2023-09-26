package main

import (
	"fmt"
)

func main() {
	// Demander à l'utilisateur le nombre de transactions
	var numTransactions int
	fmt.Println("Entrez le nombre de transactions:")
	fmt.Scan(&numTransactions)

	// Initialiser une variable pour stocker le temps total
	var totalTime float64

	// Boucler pour demander à l'utilisateur le temps de chaque transaction
	for i := 0; i < numTransactions; i++ {
		var transactionTime float64
		fmt.Printf("Entrez le temps de la transaction %d (en secondes) : ", i+1)
		fmt.Scan(&transactionTime)
		totalTime += transactionTime
	}

	// Demander à l'utilisateur l'unité de temps pour le calcul du volume (seconde, minute, heure)
	var unit string
	fmt.Println("Entrez l'unité de temps (s pour seconde, m pour minute, h pour heure):")
	fmt.Scan(&unit)

	// Calculer le volume de transactions par unité de temps
	var transactionsPerUnit float64
	switch unit {
	case "s": // par seconde
		transactionsPerUnit = float64(numTransactions) / totalTime
	case "m": // par minute
		transactionsPerUnit = float64(numTransactions) / (totalTime / 60)
	case "h": // par heure
		transactionsPerUnit = float64(numTransactions) / (totalTime / 3600)
	default:
		fmt.Println("Unité de temps non reconnue")
		return
	}

	// Afficher le résultat
	fmt.Printf("Le volume de transactions est de %.2f transactions par %s\n", transactionsPerUnit, unit)
}
