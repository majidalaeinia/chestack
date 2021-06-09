package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"log"
)

func Run() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?autocommit=true&parseTime=true&loc=%s",
		viper.GetString("DB.USERNAME"),
		viper.GetString("DB.PASSWORD"),
		viper.GetString("DB.HOST"),
		viper.GetString("DB.PORT"),
		viper.GetString("DB.NAME"),
		"Asia%2FTehran"))
	if err != nil {
		log.Println(err.Error())
		return
	}
	//defer db.Close() //TODO
	boil.SetDB(db)
}
