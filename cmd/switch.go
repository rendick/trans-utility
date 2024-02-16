package cmd

import (
	"os"

	"github.com/rendick/cts/settings"
)

func Switch() {
	if len(os.Args) < 2 {
		settings.Usage()
		os.Exit(0)
	} else {
		switch os.Args[1] {
		case "--help", "-h":
			settings.Help()
		default:
			Translate()
		}
	}
}
