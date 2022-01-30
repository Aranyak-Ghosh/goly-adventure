package jwt

import (
	"fmt"

	"github.com/Aranyak-Ghosh/spotify/utils/http"
	"github.com/golang-jwt/jwt"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type JwtManager struct {
	publicKey    []byte
	authEndpoint string
}

var authEndpoint = "https://golem-lasc-cxstg.azurewebsites.net/"

func (t JwtManager) ParseJWTtoken(tokenString string) (jwt.Token, error) {
	// Read the private key which will be used to parse the JWT
	signed, _ := jwt.ParseRSAPublicKeyFromPEM(t.publicKey)
	parsedToken, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}
		return signed, nil
	})

	return *parsedToken, err
}

func NewTokenUtil(client *http.HttpClient, logger *zap.SugaredLogger) *JwtManager {
	var resource = "/api/v1/.well-known/key"
	var keys map[string]string
	statusCode, ok, err := client.Get(authEndpoint+resource, nil).Result(&keys)

	if !ok {
		logger.Errorf("Error while getting public key from auth endpoint: %s", err)
		panic(err)
	} else {
		if statusCode >= 200 && statusCode < 300 {
			logger.Infof("Public key retrieved from auth endpoint: %s", keys["public_key"])
			return &JwtManager{[]byte(keys["publicKey"]), authEndpoint}
		} else {
			panic(err)
		}
	}
}

var Module = fx.Option(fx.Provide(NewTokenUtil))
