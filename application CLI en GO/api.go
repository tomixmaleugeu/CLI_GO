package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const url = "https://f1-companion-api-ba5k.onrender.com/drivers"

func recupererPilotes() ([]Pilote, error) {
	client := &http.Client{Timeout: 30 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("appel à l'API impossible : %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("l'API a répondu %d (200 attendu)", resp.StatusCode)
	}

	var reponse Reponse
	if err := json.NewDecoder(resp.Body).Decode(&reponse); err != nil {
		return nil, fmt.Errorf("réponse JSON invalide : %w", err)
	}

	return reponse.Drivers, nil
}
