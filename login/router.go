package login

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

// Router Router
func Router(router *gin.RouterGroup, db *sql.DB) {
	router.POST("", handlerPOST(db))
}
