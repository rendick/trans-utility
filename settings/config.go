package settings

import (
	"fmt"
)

var (
	Red   = "\033[31m"
	Bold  = "\033[1m"
	Green = "\033[32m"
	Reset = "\033[0m"
)

const (
	URL = "https://google-translation-unlimited.p.rapidapi.com/translate"
)

func Usage() {
	fmt.Println(Bold + "Usage:" + Reset + " cts <language you want to translate into> <text to translate>")
}

func Help() {
	fmt.Println("CTS uses the ISO 639 language names. If you want to find a specific language, so visit: <https://en.wikipedia.org/wiki/List_of_ISO_639_language_codes> and check" + Bold + "Set 1" + Reset)
}
