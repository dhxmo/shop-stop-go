package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pass []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	if err != nil {
		log.Fatal("Failed to generate password: ", err)
		return ""
	}

	return string(hashed)
}
