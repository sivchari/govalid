package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

const (
	oldMarkerPrefix = "// +govalid:"
	newMarkerPrefix = "//govalid:"
)

var dryRun bool

var migrateCmd = &cobra.Command{
	Use:   "migrate [paths...]",
	Short: "Migrate old marker format to new format",
	Long: `Migrate old marker format (// +govalid:) to new format (//govalid:).

If no paths are specified, the current directory is used.

Examples:
  govalid migrate                    # Migrate all .go files in current directory
  govalid migrate ./...              # Migrate all .go files recursively
  govalid migrate file.go            # Migrate a specific file
  govalid migrate --dry-run ./...    # Preview changes without modifying files`,
	RunE: func(_ *cobra.Command, args []string) error {
		paths := args
		if len(paths) == 0 {
			paths = []string{"."}
		}

		return runMigrate(paths, dryRun)
	},
}

func init() {
	migrateCmd.Flags().BoolVar(&dryRun, "dry-run", false, "preview changes without modifying files")
}

func runMigrate(paths []string, dryRun bool) error {
	files, err := collectGoFiles(paths)
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

func collectGoFiles(paths []string) ([]string, error) {
	var files []string

	for _, path := range paths {
		collected, err := collectFromPath(path)
		if err != nil {
			return nil, err
		}

		files = append(files, collected...)
	}

	return files, nil
}

func collectFromPath(path string) ([]string, error) {
	// Handle ./... pattern
	if strings.HasSuffix(path, "/...") {
		return collectRecursive(strings.TrimSuffix(path, "/..."))
	}

	info, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("failed to stat %s: %w", path, err)
	}

	if info.IsDir() {
		return collectFromDirectory(path)
	}

	if strings.HasSuffix(path, ".go") {
		return []string{path}, nil
	}

	return nil, nil
}

func collectRecursive(basePath string) ([]string, error) {
	if basePath == "" {
		basePath = "."
	}

	var files []string

	err := filepath.Walk(basePath, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(p, ".go") {
			files = append(files, p)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to walk directory %s: %w", basePath, err)
	}

	return files, nil
}

func collectFromDirectory(path string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory %s: %w", path, err)
	}

	var files []string

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".go") {
			files = append(files, filepath.Join(path, entry.Name()))
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
