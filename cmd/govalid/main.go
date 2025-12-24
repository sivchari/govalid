// Package main is the entry point for the govalid command line tool.
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	govalid_pkg "github.com/sivchari/govalid"
)

var rootCmd = &cobra.Command{
	Use:     "govalid [flags] [packages]",
	Short:   "govalid generates type-safe validation code for Go structs",
	Version: govalid_pkg.Version,
	Args:    cobra.ArbitraryArgs,
	RunE: func(_ *cobra.Command, _ []string) error {
		return runGenerator()
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
