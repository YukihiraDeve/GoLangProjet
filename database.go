package database

import (
	"log"
	// importez le package model ici
	"Golang/model"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabasePostgres struct {
	Db *gorm.DB
}

var DB DatabasePostgres

func (d *DatabasePostgres) connect() {
	connectionString := "host=host.docker.internal port=55001 user=postgres dbname=postgres password=postgrespw"
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Error : Unable to connect to database : %v", err)
	}
	log.Println("Connected to database")

	log.Println("Running migration")
	db.AutoMigrate(&model.User{})

	DB = DatabasePostgres{
		Db: db,
	}
}
