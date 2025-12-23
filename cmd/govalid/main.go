// Package main is the entry point for the govalid command line tool.
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	govalid_pkg "github.com/sivchari/govalid"
)

var rootCmd = &cobra.Command{
	Use:     "govalid",
	Short:   "govalid generates type-safe validation code for Go structs",
	Version: govalid_pkg.Version,
	Run: func(_ *cobra.Command, _ []string) {
		if err := runGenerator(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func main() {
	// Check if running in generator mode (called via go generate)
	// In this case, pass through to the generator directly
	if len(os.Args) > 1 && os.Args[1] != "migrate" && os.Args[1] != "help" && os.Args[1] != "--help" && os.Args[1] != "-h" && os.Args[1] != "version" && os.Args[1] != "--version" && os.Args[1] != "-v" {
		if err := runGenerator(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		return
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
