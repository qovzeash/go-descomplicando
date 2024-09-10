package crud

import (
	"qovzeash/go-descomplicando/crud/models"
	"qovzeash/go-descomplicando/utils"
)

func Migrate() {
	db, err := Connection()
	utils.CheckErr(err)

	db.AutoMigrate(&models.User{})
}
