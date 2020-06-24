package api

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lffranca/suflex-api/auth"
	"github.com/lffranca/suflex-api/login"
	"github.com/lffranca/suflex-api/user"
)

// Init Init
func Init(router *gin.RouterGroup, db *sql.DB) {
	auth, errAuth := auth.Init()
	if errAuth != nil {
		log.Fatalln(errAuth)
	}

	login.Router(router.Group("login"), db)

	routerPrivate := router.Group("private")
	routerPrivate.Use(auth.MiddlewareAuth())
	{
		user.RouterPrivate(routerPrivate.Group("user"), db)
	}
}
