package crud

import (
	"qovzeash/go-descomplicando/utils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("crud/crud.db"), &gorm.Config{})
	utils.CheckErr(err)

	return db, nil

}
