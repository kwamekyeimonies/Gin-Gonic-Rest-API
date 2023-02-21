package utils

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(ttl time.Duration,payload interface{},privateKey string) (string,error){

	decodedPrivateKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil{
		fmt.Println("Error decoding Private Key")
		
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if err != nil{
		fmt.Println("Error Creating Key Parse")
	}

	now := time.Now().UTC()
	claims := make(jwt.MapClaims)
	claims["sub"] = payload
	cl
}