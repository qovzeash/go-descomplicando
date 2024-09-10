package crud

import (
	"qovzeash/go-descomplicando/crud/models"
	"qovzeash/go-descomplicando/utils"
)

func CreateUser(id int, name string) {
	db, err := Connection()
	utils.CheckErr(err)

	user := models.User{
		ID:   id,
		Name: name,
	}

	db.Create(&user)
}
