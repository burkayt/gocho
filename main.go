package main

import (
	"flag"
	"github.com/spf13/viper"
	"gocho/web"
	"log"
)

func main() {
	readConfig()
	web.RegisterHandlers()
}

func readConfig() {
	var configName = flag.String("config-name", "application", "Config File Name")
	var configPath = flag.String("config-path", ".", "Config File Name")

	viper.SetConfigName(*configName)
	viper.AddConfigPath(*configPath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

}
