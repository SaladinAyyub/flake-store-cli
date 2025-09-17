package store

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/SaladinAyyub/flake-store-cli/internal/models"
)

const flakesURL = "https://raw.githubusercontent.com/SaladinAyyub/flake-store-flakes/main/flakes.json"

func FetchFlakes() ([]models.Flake, error) {
	resp, err := http.Get(flakesURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("failed to fetch flakes.json")
	}

	var flakes []models.Flake
	if err := json.NewDecoder(resp.Body).Decode(&flakes); err != nil {
		return nil, err
	}

	return flakes, nil
}
