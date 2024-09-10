package loader

import (
	"fmt"
	"os"
)

func LoadUser() {
	user := os.Getenv("USER")
	fmt.Println(user)
}
