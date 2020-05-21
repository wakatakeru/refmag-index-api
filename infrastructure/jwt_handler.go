package infrastructure

import (
	"crypto/rsa"
	"errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type JWTHandler struct {
	PubKey *rsa.PublicKey
}

type JSONWebToken struct {
	Token jwt.Token
}

type CustomClaims struct {
	Name string `json:name`
	jwt.StandardClaims
}

func NewJWTHandler() *JWTHandler {
	pubBytes, err := ioutil.ReadFile(os.Getenv("JWT_PUB_KEY_PATH"))
	if err != nil {
		log.Panic(err)
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubBytes)
	if err != nil {
		log.Panic(err)
	}

	jwtHandler := new(JWTHandler)
	jwtHandler.PubKey = pubKey

	return jwtHandler
}

// TODO: Fix(Do not bring user concept)
func (handler *JWTHandler) Verify(token string) (userName string, err error) {
	jwtHandler := NewJWTHandler()
	jsonWebToken, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtHandler.PubKey, nil
	})

	if claims, ok := jsonWebToken.Claims.(*CustomClaims); ok && jsonWebToken.Valid {
		userName = claims.Name
		return userName, nil
	}

	return userName, errors.New("Failed Parse JWT")
}
