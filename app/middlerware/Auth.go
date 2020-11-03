package middlerware

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func Auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected token!")
		}
		return []byte("secret"), nil
	})

	if token == nil || err != nil {
		result := gin.H{
			"Message": "Invalid token!",
			"Error":   err.Error(),
		}

		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
		return
	}

	fmt.Println("Token verified!")
	claims := token.Claims.(jwt.MapClaims)
	fmt.Println(claims)

	var idAccount int
	err = mapstructure.Decode(claims["account_number"], idAccount)
	if err != nil {
		result := gin.H{
			"Message": err.Error(),
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}

	c.Set("account_number", idAccount)
}
