package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/SaladinAyyub/flake-store-cli/internal/store"
	"github.com/SaladinAyyub/flake-store-cli/tui"
)

// isTerminal returns true if stdout is a terminal (interactive)
func isTerminal() bool {
	fi, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	return (fi.Mode() & os.ModeCharDevice) != 0
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all available flakes in a full TUI experience",
	Long: `fetches the latest flakes.json and list all the flakes from flake-store-flakes (SaladinAyyub/flake-store-flakes) repository and launches the full TUI.

	flake-store-cli list
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		flakes, err := store.LoadFlakesFromCache()
		if err != nil {
			fmt.Println("Cache not found, fetching from remote...")
			flakes, err = store.FetchFlakes()
			if err != nil {
				return err
			}
		}

		// Detect if terminal is interactive
		if isTerminal() {
			return tui.List(flakes)
		}

		for _, flake := range flakes {
			fmt.Printf("%s - %s\n", flake.Name, flake.Description)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
