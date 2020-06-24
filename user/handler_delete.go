package user

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lffranca/suflex-api/user_model"
)

func handlerDELETE(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idString := c.Param("id")

		id, errID := strconv.ParseInt(idString, 10, 64)
		if errID != nil {
			log.Println(errID)
			c.JSON(http.StatusBadRequest, nil)
			c.Abort()
			return
		}

		if err := user_model.ModelDelete(db, int(id)); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "Erro na requisição, tente novamente mais tarde!"})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Deletado com sucesso!"})
	}
}
