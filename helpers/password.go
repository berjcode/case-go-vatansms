package helpers

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func GenerateSalt() (string, error) {
	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}

func HashPassword(password, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(password + salt))
	hashedPassword := hex.EncodeToString(hash.Sum(nil))
	return hashedPassword
}

func CheckPassword(password, salt, hashedPassword string) bool {
	hashed := HashPassword(password, salt)
	return hashed == hashedPassword
}
