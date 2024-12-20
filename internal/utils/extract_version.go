package utils

import "strings"

func ExtractLastVersion(changelogContent string) string {
	lines := strings.Split(changelogContent, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "## [v") {
			return strings.TrimPrefix(strings.Split(line, "]")[0], "## [v")
		}
	}
	return ""
}
