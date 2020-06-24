package auth

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lffranca/suflex-api/constant"
)

// MiddlewareAuth MiddlewareAuth
func (item *Auth) MiddlewareAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := Header{}
		if err := c.ShouldBindHeader(&h); err != nil {
			log.Println(err)
			c.JSON(http.StatusForbidden, nil)
			c.Abort()
			return
		}

		resultBearer := strings.Split(h.Authorization, " ")

		if len(resultBearer) < 2 {
			c.JSON(http.StatusForbidden, nil)
			c.Abort()
			return
		}

		tokenString := resultBearer[1]

		_, errVerifyToken := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return constant.HMACSampleSecret, nil
		})
		if errVerifyToken != nil {
			log.Println(errVerifyToken)
			c.JSON(http.StatusForbidden, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
