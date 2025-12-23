package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/tools/go/packages"
)

const (
	oldMarkerPrefix = "// +govalid:"
	newMarkerPrefix = "//govalid:"
)

var dryRun bool

var migrateCmd = &cobra.Command{
	Use:   "migrate [patterns...]",
	Short: "Migrate old marker format to new format",
	Long: `Migrate old marker format (// +govalid:) to new format (//govalid:).

If no patterns are specified, the current directory is used.

Examples:
  govalid migrate                    # Migrate all .go files in current directory
  govalid migrate ./...              # Migrate all .go files recursively
  govalid migrate ./pkg/...          # Migrate files in pkg directory
  govalid migrate --dry-run ./...    # Preview changes without modifying files`,
	RunE: func(_ *cobra.Command, args []string) error {
		patterns := args
		if len(patterns) == 0 {
			patterns = []string{"."}
		}

		return runMigrate(patterns, dryRun)
	},
}

func init() {
	migrateCmd.Flags().BoolVar(&dryRun, "dry-run", false, "preview changes without modifying files")
}

func runMigrate(patterns []string, dryRun bool) error {
	files, err := collectGoFiles(patterns)
	if err != nil {
		return err
	}

	totalMigrated := 0

	for _, file := range files {
		count, err := migrateFile(file, dryRun)
		if err != nil {
			return fmt.Errorf("failed to migrate %s: %w", file, err)
		}

		totalMigrated += count
	}

	printMigrationSummary(totalMigrated, dryRun)

	return nil
}

func collectGoFiles(patterns []string) ([]string, error) {
	cfg := &packages.Config{
		Mode: packages.NeedFiles,
	}

	pkgs, err := packages.Load(cfg, patterns...)
	if err != nil {
		return nil, fmt.Errorf("failed to load packages: %w", err)
	}

	seen := make(map[string]struct{})

	var files []string

	for _, pkg := range pkgs {
		for _, file := range pkg.GoFiles {
			if _, ok := seen[file]; ok {
				continue
			}

			seen[file] = struct{}{}

			files = append(files, file)
		}
	}

	return files, nil
}

func printMigrationSummary(totalMigrated int, dryRun bool) {
	switch {
	case totalMigrated == 0:
		fmt.Println("No markers to migrate.")
	case dryRun:
		fmt.Printf("\nDry run: %d marker(s) would be migrated. Run without --dry-run to apply changes.\n", totalMigrated)
	default:
		fmt.Printf("\nMigrated %d marker(s).\n", totalMigrated)
	}
}

func migrateFile(path string, dryRun bool) (int, error) {
	cleanPath := filepath.Clean(path)

	content, err := os.ReadFile(cleanPath)
	if err != nil {
		return 0, fmt.Errorf("failed to read file: %w", err)
	}

	lines := strings.Split(string(content), "\n")
	modified := false
	count := 0

	for i, line := range lines {
		trimmed := strings.TrimLeft(line, " \t")
		if strings.HasPrefix(trimmed, oldMarkerPrefix) {
			indent := line[:len(line)-len(trimmed)]
			newLine := indent + newMarkerPrefix + strings.TrimPrefix(trimmed, oldMarkerPrefix)

			if dryRun {
				fmt.Printf("%s:%d\n", path, i+1)
				fmt.Printf("  - %s\n", line)
				fmt.Printf("  + %s\n", newLine)
			}

			lines[i] = newLine
			modified = true
			count++
		}
	}

	if modified && !dryRun {
		newContent := strings.Join(lines, "\n")

		//#nosec G306 -- file permissions match original file
		if err := os.WriteFile(cleanPath, []byte(newContent), 0o644); err != nil {
			return 0, fmt.Errorf("failed to write file: %w", err)
		}

		fmt.Printf("%s: migrated %d marker(s)\n", path, count)
	}

	return count, nil
}
