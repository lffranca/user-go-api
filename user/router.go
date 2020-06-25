package user

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

// RouterPrivate RouterPrivate
func RouterPrivate(router *gin.RouterGroup, db *sql.DB) {
	router.GET("", handlerGET(db))
	router.GET(":id", handlerByIDGET(db))
	router.POST("", handlerPOST(db))
	router.PUT(":id", handlerPUT(db))
	router.DELETE(":id", handlerDELETE(db))
}
