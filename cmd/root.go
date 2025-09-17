/*
Copyright © 2025 Saladin <dev@saladin.pro>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "flake-store-cli",
	Short: "Search, Install & Use Nix flakes in seconds",
	Long: `flake-store-cli is a simple tool to browse and use Nix flakes 
from the flake-store-flakes repository.

You can:
  • list all available flakes and launch the complete TUI experience
  • search by name or description in CLI mode
  • quickly set up your current directory to use a chosen flake 
    (copies flake.nix and generates a .envrc to work with direnv)

This makes it easy to setup your development workflow using nix flakes.
Lets go type ->

	flake-store-cli list
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.flake-store-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
