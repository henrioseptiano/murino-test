package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func main() {

}

func generateJWT(userID string) (*string, error) {
	secretKey := "murino"
	token := jwt.New(jwt.SigningMethodHS256)

	// Create claims
	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = "your_issuer"
	claims["sub"] = "your_subject"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Expires in 24 hours
	claims["iat"] = time.Now().Unix()
	claims["user_id"] = userID

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println("Error signing token:", err)
		return nil, err
	}

	return &tokenString, nil
}

// hmacSHA256 calculates the HMAC with SHA-256 hash of the given data and key.
func hmacSHA256(data string, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write([]byte(data))
	return h.Sum(nil)
}
