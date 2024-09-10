package crud

import (
	"fmt"
	"qovzeash/go-descomplicando/crud/models"
	"qovzeash/go-descomplicando/utils"
)

func ListUser(id int) {
	db, err := Connection()
	utils.CheckErr(err)

	user := models.User{
		ID: id,
	}

	db.First(&user)

	fmt.Println(user)

}
