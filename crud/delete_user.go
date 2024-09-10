package crud

import (
	"qovzeash/go-descomplicando/crud/models"
	"qovzeash/go-descomplicando/utils"
)

func DeleteUser(id int) {
	db, err := Connection()
	utils.CheckErr(err)

	user := models.User{
		ID: id,
	}

	db.Delete(&user)
}
