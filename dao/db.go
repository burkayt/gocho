package dao

import (
	"bytes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"log"
)

type DbConfig struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Schema   string `mapstructure:"schema"`
}

type SqlxDatabase struct {
	db     *sqlx.DB
	config DbConfig
}

type Database interface {
	Connect() *sqlx.DB
	GetConnection() *sqlx.DB
}

func NewSqlxDatabase(config DbConfig) Database {
	database := SqlxDatabase{config: config}
	database.db = database.Connect()
	return &database

}

func (database SqlxDatabase) GetConnection() *sqlx.DB {
	return database.db
}

func (database SqlxDatabase) Connect() *sqlx.DB {
	datasource := database.config.getDatasource()

	db, err := sqlx.Connect("mysql", datasource)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}

func GetDbConfig() DbConfig {
	config := DbConfig{}
	err := viper.UnmarshalKey("db", &config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}

func (config DbConfig) getDatasource() string {
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
