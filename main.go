package main

import (
	"flag"
	"fmt"

	"github.com/taua-almeida/changelog-forge/cmd"
)

func main() {
	generateJson := flag.Bool("generate-json", false, "Generate changelog input JSON")
	updateChangelog := flag.Bool("update-changelog", false, "Update the CHANGELOG.md file")
	help := flag.Bool("help", false, "Show help")

	flag.Parse()

	if *help {
		fmt.Println("Available options:")
		fmt.Println("--generate-json: Generate changelog input JSON file")
		fmt.Println("--update-changelog: Update the CHANGELOG.md file")
		return
	} else if *generateJson {
		cmd.GenerateJSON()
	} else if *updateChangelog {
		cmd.UpdateChangelog()
	} else {
		fmt.Println("Usage: changelog-forge --generate-json or --update-changelog")
	}

}
