package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/rendick/cts/settings"
)

func Translate() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cts <language you want to translate into> <text to translate>")
		os.Exit(0)
	} else {

		var words []string
		for _, arg := range os.Args[2:] {
			words = append(words, arg)
		}

		Url = fmt.Sprintf("texte=%s&to_lang=%s", strings.Join(words, " "), os.Args[1])
		payload := strings.NewReader(Url)

		req, _ := http.NewRequest("POST", settings.URL, payload)

		req.Header.Add("content-type", "application/x-www-form-urlencoded")
		req.Header.Add("X-RapidAPI-Key", settings.API_KEY)
		req.Header.Add("X-RapidAPI-Host", "google-translation-unlimited.p.rapidapi.com")

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)

		var apiResponse APIResponse

		err := json.Unmarshal(body, &apiResponse)
		if err != nil {
			log.Fatal("Error!")
			os.Exit(0)
		}

		fmt.Printf("Original text: %s\nTranslated text: %s\n", apiResponse.TranslationData.Original, apiResponse.TranslationData.Translation)
	}
}
