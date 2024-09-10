package main

import "qovzeash/go-descomplicando/crud"

func main() {
	crud.Migrate()
	crud.CreateUser(1, "Élodie Durand")
	crud.ListUser(1)
	crud.UpdateUser(1, "Julien Lefèvre")
	crud.ListUser(1)
}
