package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(getPassword string) string {
	salt := 8
	password := []byte(getPassword)
	hash, err := bcrypt.GenerateFromPassword(password, salt)
	if err != nil {
		panic(err)
	}

	return string(hash)
}

func ComparePassword(hashPassword, getPassword []byte) bool {
	password := []byte(getPassword)
	hash := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(hash, password)

	return err == nil
}
