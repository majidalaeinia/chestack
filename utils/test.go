package utils

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-txdb"
	"github.com/MajidAlaeinia/chestack/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"log"
	"time"
)

func RunDatabase() *sqlx.DB {
	config.ReadConfig()
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?autocommit=true&parseTime=true&loc=%s",
		viper.GetString("DB.USERNAME"),
		viper.GetString("DB.PASSWORD"),
		viper.GetString("DB.HOST"),
		viper.GetString("DB.PORT"),
		viper.GetString("DB.NAME"),
		"Asia%2FTehran"))
	if err != nil {
		log.Println(fmt.Sprintf("Error: %v", err))
	}

	boil.SetDB(db)

	db.SetMaxOpenConns(viper.GetInt("DB.MAX_OPEN_CONNS"))
	db.SetMaxIdleConns(viper.GetInt("DB.MAX_IDLE_CONNS"))
	db.SetConnMaxLifetime(viper.GetDuration("DB.CONN_MAX_LIFE_TIME") * time.Minute)

	return db
}

func RunTxDb() *sql.DB {
	config.ReadConfig()
	txdb.Register("txdb", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?autocommit=true&parseTime=true&loc=%s",
		viper.GetString("DB.USERNAME"),
		viper.GetString("DB.PASSWORD"),
		viper.GetString("DB.HOST"),
		viper.GetString("DB.PORT"),
		viper.GetString("DB.NAME"),
		"Asia%2FTehran"))
	db, err := sql.Open("txdb", "identifier")
	if err != nil {
		log.Fatal(err)
	}
	boil.SetDB(db)

	return db
}
