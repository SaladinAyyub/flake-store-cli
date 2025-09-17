package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/SaladinAyyub/flake-store-cli/internal/store"
)

var searchCmd = &cobra.Command{
	Use:   "search [keyword]",
	Short: "Search flakes by name or description",
	Long: `Searches the flakes from flake-store-flakes repository 
by matching the given keyword in the name or description.

Example:
  flake-store-cli search raylib
`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		keyword := strings.ToLower(args[0])

		// Load flakes (cache first)
		flakes, err := store.LoadFlakesFromCache()
		if err != nil {
			fmt.Println("Cache not found, fetching from remote...")
			flakes, err = store.FetchFlakes()
			if err != nil {
				return err
			}
		}

		found := false
		for _, flake := range flakes {
			if strings.Contains(strings.ToLower(flake.Name), keyword) ||
				strings.Contains(strings.ToLower(flake.Description), keyword) {
				fmt.Printf("%s - %s\n%s\n\n", flake.Name, flake.Description, flake.RepoURL)
				found = true
			}
		}

		if !found {
			fmt.Println("No flakes found matching:", keyword)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
