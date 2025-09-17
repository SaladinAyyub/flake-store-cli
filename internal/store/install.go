package store

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// InstallFlake installs the specified flake into the current directory.
// nodirenv = true skips creating .envrc
func InstallFlake(flakeName string, nodirenv bool) error {
	// Load flakes from cache (or fetch from remote)
	flakes, err := LoadFlakesFromCache()
	if err != nil {
		fmt.Println("Cache not found, fetching from remote...")
		flakes, err = FetchFlakes()
		if err != nil {
			return err
		}
	}

	// Find the requested flake
	var repoURL string
	found := false
	for _, f := range flakes {
		if f.Name == flakeName {
			repoURL = f.RepoURL
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("flake %q not found", flakeName)
	}

	// Determine current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	// GitHub raw content URL
	rawURL := strings.Replace(
		repoURL,
		"https://github.com/",
		"https://raw.githubusercontent.com/",
		1,
	)
	rawURL = strings.Replace(rawURL, "/tree/", "/", 1)

	// Files to copy
	files := []string{"flake.nix", "shell.nix"}

	for _, file := range files {
		fileURL := fmt.Sprintf("%s/%s", rawURL, file)
		resp, err := http.Get(fileURL)
		if err != nil {
			return fmt.Errorf("failed to download %s: %w", file, err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == 404 {
			if file == "shell.nix" {
				// optional file, skip if missing
				continue
			} else {
				return fmt.Errorf("file %s not found (status %d)", file, resp.StatusCode)
			}
		} else if resp.StatusCode != 200 {
			return fmt.Errorf("failed to download %s: status %d", file, resp.StatusCode)
		}

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		destPath := filepath.Join(cwd, file)
		if err := os.WriteFile(destPath, data, 0o644); err != nil {
			return fmt.Errorf("failed to write file %s: %w", destPath, err)
		}
	}

	// Create .envrc unless nodirenv
	if !nodirenv {
		envrcPath := filepath.Join(cwd, ".envrc")
		if _, err := os.Stat(envrcPath); os.IsNotExist(err) {
			if err := os.WriteFile(envrcPath, []byte("use flake\n"), 0o644); err != nil {
				return fmt.Errorf("failed to create .envrc: %w", err)
			}
		}
	}

	fmt.Printf("Successfully installed flake %q in %s\n", flakeName, cwd)
	return nil
}
