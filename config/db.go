package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectDB() (*gorm.DB, error) {
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := gorm.Open("postgres", config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}

	// Set up connection pool
	conn.DB().SetMaxIdleConns(20)
	conn.DB().SetMaxOpenConns(200)

	return conn, err
}
