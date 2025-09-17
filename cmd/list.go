package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/SaladinAyyub/flake-store-cli/internal/store"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all available flakes.",
	Long: `fetches the latest flakes.json and list all the flakes from flake-store-flakes (SaladinAyyub/flake-store-flakes) repository

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

		for _, flake := range flakes {
			fmt.Printf("%s - %s\n", flake.Name, flake.Description)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
