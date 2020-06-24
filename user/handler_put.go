package user

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handlerPUT(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	}
}
