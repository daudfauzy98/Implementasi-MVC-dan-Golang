package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashGenerator untuk merubah password menjadi kode hash
func HashGenerator(str string) (string, error) {
	hashedString, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedString), nil
}

// HashComparator untuk membandingkan hashed password dari inputan user dengan yg ada di database
func HashComparator(hashedByte []byte, passwordByte []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedByte, passwordByte)
	if err != nil {
		return err
	}
	return nil
}
