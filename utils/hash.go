package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(p string) (string, error) {
	hp, err := bcrypt.GenerateFromPassword([]byte(p), 14)

	return string(hp), err
}

func CheckPasswordHash(hp, p string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hp), []byte(p))
	return err == nil
}
