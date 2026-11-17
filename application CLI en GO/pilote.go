package main

import "strconv"

type Pilote struct {
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Nationality string `json:"nationality"`
	Number      string `json:"number"`
	Code        string `json:"code"`
}

func (p Pilote) NomComplet() string {
	return p.FirstName + " " + p.LastName
}

type Reponse struct {
	Drivers []Pilote `json:"drivers"`
}

func chercherPilote(pilotes []Pilote, numero int) (Pilote, bool) {
	cible := strconv.Itoa(numero)

	for _, pilote := range pilotes {
		if pilote.Number == cible {
			return pilote, true
		}
	}

	return Pilote{}, false
}
