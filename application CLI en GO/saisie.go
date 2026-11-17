package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const tousLesPilotes = -1

func demanderNumero() (string, error) {
	fmt.Print("Entrez un numéro de pilote : ")

	lecteur := bufio.NewScanner(os.Stdin)
	if !lecteur.Scan() {
		if err := lecteur.Err(); err != nil {
			return "", fmt.Errorf("lecture de la saisie impossible : %w", err)
		}
		return "", errors.New("aucune saisie")
	}

	return strings.TrimSpace(lecteur.Text()), nil
}

func verifierNumero(saisie string) (int, error) {
	if strings.EqualFold(saisie, "all") {
		return tousLesPilotes, nil
	}

	numero, err := strconv.Atoi(saisie)
	if err != nil || numero <= 0 {
		return 0, fmt.Errorf("%q n'est pas un numéro valide.", saisie)
	}

	return numero, nil
}
