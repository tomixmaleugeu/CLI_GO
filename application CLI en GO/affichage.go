package main

import "fmt"

func afficherPilote(pilote Pilote) {
	fmt.Println()
	fmt.Println("--- Pilote trouvé ---")
	fmt.Printf("Numéro      : %s\n", pilote.Number)
	fmt.Printf("Nom         : %s\n", pilote.NomComplet())
	fmt.Printf("Nationalité : %s\n", pilote.Nationality)
	fmt.Printf("Code        : %s\n", pilote.Code)
}

func afficherTous(pilotes []Pilote) {
	fmt.Printf("\n--- %d pilotes ---\n", len(pilotes))

	for _, pilote := range pilotes {
		fmt.Printf("%-4s %-22s %-14s %s\n", pilote.Number, pilote.NomComplet(), pilote.Nationality, pilote.Code)
	}
}
