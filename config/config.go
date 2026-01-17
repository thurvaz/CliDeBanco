package config

import "gorm.io/gorm"

var (
	db *gorm.DB
)

func Init() error {
	var err error

	db, err = InitializeSQLite()

	if err != nil {
		panic(err)
	}
	return nil
}

func GetSQLite() *gorm.DB {
	return db
}
