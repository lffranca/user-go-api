package user

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handlerPOST(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body UserRequest
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Parametros invalidos"})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, body)
	}
}
