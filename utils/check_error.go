package utils

import "log"

func CheckErr(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
