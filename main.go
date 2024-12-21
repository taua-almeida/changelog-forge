package main

import (
	"bytes"
	"flag"
	"fmt"
	"os/exec"
	"strings"

	"github.com/taua-almeida/changelog-forge/cmd"
)

var Version = "unknown"

func getGitVersion() string {
	cmd := exec.Command("git", "describe", "--tags", "--abbrev=0")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "unknown"
	}
	return strings.TrimSpace(out.String())
}

func main() {
	generateJson := flag.Bool("generate-json", false, "Generate changelog input JSON")
	updateChangelog := flag.Bool("update-changelog", false, "Update the CHANGELOG.md file")
	showVersion := flag.Bool("version", false, "Show the current version")

	help := flag.Bool("help", false, "Show help")

	flag.Parse()

	// Fallback to Git tag if Version is unknown
	if Version == "unknown" {
		Version = getGitVersion()
	}

	if *help {
		fmt.Println("Available options:")
		fmt.Println("--generate-json: Generate changelog input JSON file")
		fmt.Println("--update-changelog: Update the CHANGELOG.md file")
		fmt.Println("--version: Show the current version")
		return
	} else if *generateJson {
		cmd.GenerateJSON()
	} else if *updateChangelog {
		newVersion := cmd.UpdateChangelog()
		fmt.Println(newVersion)
	} else if *showVersion {
		fmt.Println("Version:", Version)
	} else {
		fmt.Println("Usage: changelog-forge --generate-json or --update-changelog")
	}

}
