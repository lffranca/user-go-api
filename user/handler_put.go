package user

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lffranca/suflex-api/hash"
	"github.com/lffranca/suflex-api/token"
	"github.com/lffranca/suflex-api/user_model"
)

func handlerPUT(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idString := c.Param("id")

		id, errID := strconv.ParseInt(idString, 10, 64)
		if errID != nil {
			log.Println(errID)
			c.JSON(http.StatusBadRequest, nil)
			c.Abort()
			return
		}

		var body UserRequest
		if err := c.ShouldBindJSON(&body); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "Parametros invalidos"})
			c.Abort()
			return
		}

		salt := token.GenerateToken(20)

		pass := fmt.Sprintf("%s%s", body.Senha, salt)
		hash, errHash := hash.GenerateHash(pass)
		if errHash != nil {
			log.Println(errHash)
			c.JSON(http.StatusBadRequest, gin.H{"message": "Erro na requisição, tente novamente mais tarde!"})
			c.Abort()
			return
		}

		if err := user_model.ModelUpdate(
			db,
			int(id),
			body.Nome,
			body.Sobrenome,
			body.Username,
			hash,
			salt,
		); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "Erro na requisição, tente novamente mais tarde!"})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Alterado com sucesso!"})
	}
}
