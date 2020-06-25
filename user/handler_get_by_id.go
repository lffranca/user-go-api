package user

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lffranca/suflex-api/user_model"
)

func handlerByIDGET(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idString := c.Param("id")

		id, errID := strconv.ParseInt(idString, 10, 64)
		if errID != nil {
			log.Println(errID)
			c.JSON(http.StatusBadRequest, nil)
			c.Abort()
			return
		}

		item, errItem := user_model.ModelGetByID(db, int(id))
		if errItem != nil {
			log.Println(errItem)
			c.JSON(http.StatusBadRequest, gin.H{"message": "Erro na requisição, tente novamente mais tarde!"})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, item)
	}
}
