package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/taua-almeida/changelog-forge/internal/versioning"
)

func UpdateChangelog() string {
	inputFile := "changelog.json"
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("failed to read file: %v\n", err)
		os.Exit(1)
	}

	var entry struct {
		Version      string   `json:"version"`
		Date         string   `json:"date"`
		Descriptions []string `json:"descriptions"`
	}

	if err := json.Unmarshal(data, &entry); err != nil {
		fmt.Printf("failed to unmarshal JSON: %v\n", err)
		os.Exit(1)
	}

	changelogData, err := os.ReadFile("CHANGELOG.md")
	if err != nil {
		fmt.Printf("failed to read CHANGELOG.md: %v\n", err)
		os.Exit(1)
	}

	// Extract the last version
	lastVersion := extractLastVersion(string(changelogData))
	if lastVersion == "" {
		fmt.Println("failed to extract last version")
		os.Exit(1)
	}

	// Increment version
	newVersion, err := versioning.IncrementVersion(lastVersion, entry.Version)
	if err != nil {
		fmt.Printf("failed to increment version: %v\n", err)
		os.Exit(1)
	}

	// Format date
	date := entry.Date
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}

	// Format descriptions
	descriptions := ""
	for _, desc := range entry.Descriptions {
		trimmedDesc := strings.TrimSpace(desc)
		capitalizedDesc := strings.ToUpper(string(trimmedDesc[0])) + trimmedDesc[1:]
		descriptions += fmt.Sprintf("- %s\n", capitalizedDesc)
	}

	// Create the new changelog entry
	newEntry := fmt.Sprintf("## [v%s] - %s\n%s", newVersion, date, descriptions)

	// Insert the new entry after the "All changes to this project..." header
	updatedChangelog := insertNewEntry(string(changelogData), newEntry)

	// Write updated changelog back to the file
	err = os.WriteFile("CHANGELOG.md", []byte(updatedChangelog), 0644)
	if err != nil {
		fmt.Printf("failed to write updated changelog: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("CHANGELOG.md updated successfully!")
	return newVersion
}

func extractLastVersion(changelogContent string) string {
	lines := strings.Split(changelogContent, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "## [v") {
			return strings.TrimPrefix(strings.Split(line, "]")[0], "## [v")
		}
	}
	return ""
}

func insertNewEntry(changelogContent, newEntry string) string {
	lines := strings.Split(changelogContent, "\n")
	for i, line := range lines {
		if strings.Contains(line, "All changes to this project") {
			// Insert new entry after the header
			return strings.Join(lines[:i+2], "\n") + "\n\n" + newEntry + strings.Join(lines[i+2:], "\n")
		}
	}
	// If the header is not found, append to the top
	return newEntry + "\n\n" + changelogContent
}
