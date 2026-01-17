package config

import (
	"os"

	"github.com/glebarez/sqlite"
	"github.com/thurvaz/CliDeBanco/schemas"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}

		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		panic(err)
	}

	return db, nil
}
