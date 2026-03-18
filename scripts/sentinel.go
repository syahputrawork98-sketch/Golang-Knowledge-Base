package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("--- Sentinel Audit (Go Edition) ---")
	basePath, _ := os.Getwd()
	fmt.Printf("Auditing: %s\n\n", basePath)

	errors := []string{}

	// 1. Check Standards
	standardsFiles := []string{
		"architecture.md", "conventions.md", "workflow.md",
		"status-protocol.md", "contribution.md", "core-contribution.md",
	}
	for _, f := range standardsFiles {
		path := filepath.Join(basePath, "docs", "standards", f)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			errors = append(errors, fmt.Sprintf("Missing standard file: %s", f))
		}
	}

	// 2. Audit Structure
	err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			name := info.Name()
			if strings.HasPrefix(name, ".") || name == "scripts" || name == "docs" {
				if name != "." && name != "docs" {
					return filepath.SkipDir
				}
			}

			if strings.HasPrefix(name, "RAK-") || strings.HasPrefix(name, "SR-") ||
				strings.HasPrefix(name, "BK-") || strings.HasPrefix(name, "CH-") {
				
				readmePath := filepath.Join(path, "README.md")
				if _, err := os.Stat(readmePath); os.IsNotExist(err) {
					errors = append(errors, fmt.Sprintf("Missing README.md in %s", path))
				}

				if strings.HasPrefix(name, "CH-") {
					assets := filepath.Join(path, "assets")
					if _, err := os.Stat(assets); os.IsNotExist(err) {
						errors = append(errors, fmt.Sprintf("Missing 'assets/' folder in %s", path))
					}
					examples := filepath.Join(path, "examples")
					if _, err := os.Stat(examples); os.IsNotExist(err) {
						errors = append(errors, fmt.Sprintf("Missing 'examples/' folder in %s", path))
					}
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Walking error: %v\n", err)
	}

	if len(errors) == 0 {
		fmt.Println("[PASS] Everything is perfectly standardized! (Gold Standard)")
	} else {
		fmt.Printf("[FAIL] Found %d inconsistencies:\n", len(errors))
		for _, e := range errors {
			fmt.Printf(" - %s\n", e)
		}
		os.Exit(1)
	}
}
