package cmd

import (
	"os"

	"github.com/rendick/trans-utility/settings"
)

func Switch() {
	if len(os.Args) < 2 {
		settings.Usage()
		os.Exit(0)
	} else {
		switch os.Args[1] {
		case "--help", "-h":
			settings.Help()
		case "--version", "-v":
			settings.Version()
		default:
			Translate()
		}
	}
}
