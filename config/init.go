package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path"
	"runtime"
)

func ReadConfig() {
	env := "development"
	if os.Getenv("ENVIRONMENT") == "production" {
		env = "production"
	}

	viper.SetConfigName(env)
	viper.SetConfigType("toml")
	viper.AddConfigPath(configPath())
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func configPath() string {
	_, fullFilename, _, _ := runtime.Caller(0)
	pwd := path.Dir(fullFilename)
	return pwd
}
