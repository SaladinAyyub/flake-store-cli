package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/SaladinAyyub/flake-store-cli/internal/store"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch the latest flakes from the flake-store-flakes repository",
	Long: `Fetches flakes.json from the flake-store-flakes repo on GitHub 
and updates the local cache so it can be reused by other commands.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		flakes, err := store.FetchFlakes()
		if err != nil {
			return err
		}

		fmt.Printf("Fetched %d flakes and updated cache.\n", len(flakes))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}
