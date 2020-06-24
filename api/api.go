package api

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lffranca/suflex-api/auth"
	"github.com/lffranca/suflex-api/user"
)

// Init Init
func Init(router *gin.RouterGroup, db *sql.DB) {
	auth, errAuth := auth.Init()
	if errAuth != nil {
		log.Fatalln(errAuth)
	}

	routerPrivate := router.Group("private")
	routerPrivate.Use(auth.MiddlewareAuth())
	{
		user.RouterPrivate(routerPrivate, db)
	}
}
