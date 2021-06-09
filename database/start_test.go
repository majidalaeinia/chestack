package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	readConfig()
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?autocommit=true&parseTime=true&loc=%s",
		viper.GetString("DB.USERNAME"),
		viper.GetString("DB.PASSWORD"),
		viper.GetString("DB.HOST"),
		viper.GetString("DB.PORT"),
		viper.GetString("DB.NAME"),
		"Asia%2FTehran"))

	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(fmt.Sprintf("Error: %v", err))
		}
	}(db)
	err = db.Ping()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func readConfig() {
	env := "development"
	if os.Getenv("ENVIRONMENT") == "production" {
		env = "production"
	}

	viper.SetConfigName(env)
	viper.SetConfigType("toml")
	viper.AddConfigPath("./../config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
