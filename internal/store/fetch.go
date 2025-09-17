package store

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/SaladinAyyub/flake-store-cli/internal/models"
)

const flakesURL = "https://raw.githubusercontent.com/SaladinAyyub/flake-store-flakes/main/flakes.json"

// getCacheFile returns the path to the cache file.
func getCacheFile() (string, error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(cacheDir, "flake-store", "flakes.json"), nil
}

// FetchFlakes downloads flakes.json and saves it to cache.
func FetchFlakes() ([]models.Flake, error) {
	resp, err := http.Get(flakesURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch flakes.json: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var flakes []models.Flake
	if err := json.Unmarshal(data, &flakes); err != nil {
		return nil, fmt.Errorf("failed to unmarshal flakes.json: %w", err)
	}

	// Cache the result
	cacheFile, err := getCacheFile()
	if err == nil {
		_ = os.MkdirAll(filepath.Dir(cacheFile), 0o755)
		_ = os.WriteFile(cacheFile, data, 0o644)
	}

	return flakes, nil
}

// LoadFlakesFromCache loads flakes.json from cache if available.
func LoadFlakesFromCache() ([]models.Flake, error) {
	cacheFile, err := getCacheFile()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(cacheFile)
	if err != nil {
		return nil, err
	}

	var flakes []models.Flake
	if err := json.Unmarshal(data, &flakes); err != nil {
		return nil, err
	}

	return flakes, nil
}
