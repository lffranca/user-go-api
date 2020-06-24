package login

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lffranca/suflex-api/constant"
	"github.com/lffranca/suflex-api/hash"
	"github.com/lffranca/suflex-api/user_model"
)

func handlerPOST(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body LoginRequest
		if err := c.ShouldBindJSON(&body); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "Parametros invalidos"})
			c.Abort()
			return
		}

		userItem, errUserItem := user_model.ModelGetByUsername(db, body.Username)
		if errUserItem != nil {
			log.Println(errUserItem)
			c.JSON(http.StatusBadRequest, gin.H{"message": "Usuário inválido"})
			c.Abort()
			return
		}

		if !hash.CheckPasswordHash(fmt.Sprintf("%s%s", body.Senha, userItem.Salt), userItem.Senha) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Usuário inválido"})
			c.Abort()
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"exp":      time.Now().Add(time.Hour * 3).UTC().Unix(),
			"username": userItem.Username,
			"name":     userItem.Nome,
			"lastname": userItem.Sobrenome,
		})

		tokenString, errTokenString := token.SignedString(constant.HMACSampleSecret)
		if errTokenString != nil {
			log.Println(errTokenString)
			c.JSON(http.StatusBadRequest, gin.H{"message": "Usuário inválido"})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	}
}
