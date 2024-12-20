package tests

import (
	"os"
	"strings"
	"testing"

	"github.com/taua-almeida/changelog-forge/cmd"
)

func TestUpdateChangelog(t *testing.T) {
	// Create a temporary changelog.json
	changelogContent := `{
		"version": "patch",
		"date": "2024-12-20",
		"descriptions": [
			"add new show",
			"fix layout issue",
			"remove unused icons"
		]
	}`
	err := os.WriteFile("changelog.json", []byte(changelogContent), 0644)
	if err != nil {
		t.Fatalf("failed to create changelog.json: %v", err)
	}

	// Create a temporary CHANGELOG.md
	initialChangelog := `# Changelog

All changes to this project will be documented in this file.

## [v1.8.21] - 2024-12-19
- Fix minor bug
`
	err = os.WriteFile("CHANGELOG.md", []byte(initialChangelog), 0644)
	if err != nil {
		t.Fatalf("failed to create CHANGELOG.md: %v", err)
	}

	// Run the UpdateChangelog function
	cmd.UpdateChangelog()

	// Verify the output in CHANGELOG.md
	updatedChangelog, err := os.ReadFile("CHANGELOG.md")
	if err != nil {
		t.Fatalf("failed to read updated CHANGELOG.md: %v", err)
	}

	expected := `## [v1.8.22] - 2024-12-20
- Add new show
- Fix layout issue
- Remove unused icons`

	if !contains(string(updatedChangelog), expected) {
		t.Errorf("expected changelog to contain:\n%s\n\nGot:\n%s", expected, string(updatedChangelog))
	}

	// Cleanup
	_ = os.Remove("changelog.json")
	_ = os.Remove("CHANGELOG.md")
}

func contains(haystack, needle string) bool {
	return strings.Contains(haystack, needle)
}
