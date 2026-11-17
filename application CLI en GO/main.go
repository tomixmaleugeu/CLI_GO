package main

import "fmt"

func main() {
	saisie, err := demanderNumero()
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	numero, err := verifierNumero(saisie)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	pilotes, err := recupererPilotes()
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	if numero == tousLesPilotes {
		afficherTous(pilotes)
		return
	}

	pilote, trouve := chercherPilote(pilotes, numero)
	if !trouve {
		fmt.Printf("Aucun pilote ne porte le numéro %d.\n", numero)
		return
	}

	afficherPilote(pilote)
}
