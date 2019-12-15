package dao

import (
	"bytes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"log"
)

type dbConfig struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Schema   string `mapstructure:"schema"`
}

var db *sqlx.DB

func InitDb() {
	config := dbConfig{}
	err := viper.UnmarshalKey("db", &config)
	if err != nil {
		log.Fatal(err)
	}

	datasource := config.getDatasource()

	db, err = sqlx.Connect("mysql", datasource)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func (config dbConfig) getDatasource() string {
	var datasource bytes.Buffer
	datasource.WriteString(config.User)
	datasource.WriteString(":")
	datasource.WriteString(config.Password)
	datasource.WriteString("@tcp(")
	datasource.WriteString(config.Host)
	datasource.WriteString(")/")
	datasource.WriteString(config.Schema)

	return datasource.String()
}
