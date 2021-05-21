package public

import (
	providers2 "github.com/MajidAlaeinia/chestack/app/providers"
	"log"
	"os"
)

func Run() {
	r := providers2.Routes()

	err := r.Run("0.0.0.0:8081")
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
