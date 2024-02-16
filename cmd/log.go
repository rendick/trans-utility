package cmd

import (
	"log"
	"os"
	"time"

	"github.com/rendick/trans-utility/settings"
)

func WriteLogs() {
	_, err := os.Stat(settings.LogPath)
	if os.IsNotExist(err) {
		err := os.Mkdir(settings.LogPath, 0750)
		if err != nil {
			log.Fatal(err)
		}
	}
	f, err := os.OpenFile(settings.LogPath+"transutil.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, writeErr := f.WriteString(time.Now().Format("2006-01-02 15:04:05") + "\n" + Log + "\n")
	if writeErr != nil {
		log.Fatal(writeErr)
	}
}

// sudo chmod 777 /var/log/transutil
