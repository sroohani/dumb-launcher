package db

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func Load(dbUrl string) {
	db, err := gorm.Open(sqlite.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening the database: %v", err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatalf("Error accessing the database object: %v", err)
	}
	defer sqlDb.Close()
}

func Save(dbUrl string) {
	db, err := gorm.Open(sqlite.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening the database: %v", err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatalf("Error accessing the database object: %v", err)
	}
	defer sqlDb.Close()
}
