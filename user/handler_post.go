package user

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handlerPOST(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	}
}
