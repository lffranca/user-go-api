package auth

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
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

		token := resultBearer[1]

		log.Println("VERIFY TOKEN HERE", token)

		c.Next()
	}
}
