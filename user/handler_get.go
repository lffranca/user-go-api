package user

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lffranca/suflex-api/user_model"
)

func handlerGET(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		items, errItems := user_model.ModelGetAll(db)
		if errItems != nil {
			log.Println(errItems)
			c.JSON(http.StatusBadRequest, gin.H{"message": "Erro na requisição, tente novamente mais tarde!"})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, items)
	}
}
