package public

import (
	"github.com/MajidAlaeinia/chestack/app/providers"
	"github.com/MajidAlaeinia/chestack/database"
	"github.com/spf13/viper"
	"log"
	"os"
)

func Run() {
	database.Run()
	r := providers.Routes()

	err := r.Run(viper.GetString("GIN.ADDRESS"))
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
