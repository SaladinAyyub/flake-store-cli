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
	Run: func(cmd *cobra.Command, args []string) {
		flakes, err := store.FetchFlakes()
		if err != nil {
			fmt.Println("❌ Error:", err)
			return
		}

		fmt.Println("Available flakes:")
		for _, f := range flakes {
			fmt.Printf("  - %s: %s\n", f.Name, f.Description)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
