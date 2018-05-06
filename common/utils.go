package utils

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	Config "gitlab.com/anant-sharma/go-boilerplate/config"
)

/*
	Get Application Configuration
*/
var config = Config.GetConfig()

var myJwtSigningKey = []byte(config.Jwt.Secret)

/*
GenToken - A Util function to generate jwtToken which can be used in the request header
*/
func GenToken(id uint) (string, error) {

	/* Create Token */
	jwtToken := jwt.New(jwt.SigningMethodHS256)

	/* Set token claims */
	jwtToken.Claims = jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Second * config.Jwt.ExpiresIn).Unix(),
		"iss": config.Jwt.Issuer,
	}

	/* Sign the token with our secret */
	return jwtToken.SignedString(myJwtSigningKey)

}
