package cmd

import (
	"github.com/spf13/cobra"

	"github.com/SaladinAyyub/flake-store-cli/internal/store"
)

var nodirenv bool

var installCmd = &cobra.Command{
	Use:   "install [flake-name]",
	Short: "Install a flake into the current directory",
	Long: `Downloads flake.nix & shell.nix from the specified flake into the current working directory.
By default, it also creates a .envrc file for direnv with "use flake".
Use --nodirenv to skip .envrc creation.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		flakeName := args[0]
		return store.InstallFlake(flakeName, nodirenv)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.Flags().BoolVar(&nodirenv, "nodirenv", false, "Do not create .envrc file")
}
