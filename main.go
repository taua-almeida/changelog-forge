package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/taua-almeida/changelog-forge/cmd"
	"github.com/taua-almeida/changelog-forge/internal/utils"
)

var Version = "unknown"

func getVersionFromChangeLog() string {
	changelogData, err := os.ReadFile("CHANGELOG.md")
	if err != nil {
		return "unknown"
	}

	return utils.ExtractLastVersion(string(changelogData))
}

func main() {
	generateJson := flag.Bool("generate-json", false, "Generate changelog input JSON")
	updateChangelog := flag.Bool("update-changelog", false, "Update the CHANGELOG.md file")
	showVersion := flag.Bool("version", false, "Show the current version")

	help := flag.Bool("help", false, "Show help")

	flag.Parse()

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
		if Version == "unknown" {
			Version = getVersionFromChangeLog()
		}
		fmt.Println("Version:", Version)
	} else {
		fmt.Println("Usage: changelog-forge --generate-json or --update-changelog")
	}

}
