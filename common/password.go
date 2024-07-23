package common

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	pw := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pw, 14)
	if err != nil {
		Logger("ERROR", "Error hashing password", err)
	}
	return string(hash)
}

func ComparePassword(password string, hash []byte) error {
	return bcrypt.CompareHashAndPassword(hash, []byte(password))
}
