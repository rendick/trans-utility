package cmd

var Url string

type TranslationData struct {
	Original    string `json:"original_text"`
	Translation string `json:"translation"`
}

type APIResponse struct {
	Meta            map[string]string `json:"meta"`
	TranslationData TranslationData   `json:"translation_data"`
}
