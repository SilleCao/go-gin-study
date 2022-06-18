package middlewares

import (
	"log"
	"net/http"

	"github.com/EDDYCJY/go-gin-example/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]

		token, err := service.NewJWTService().ValidateToken(tokenString)

		if err != nil {
			log.Println(err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[Name]", claims["name"])
			log.Println("Claims[Admin]", claims["admin"])
			log.Println("Claims[Issuer]", claims["iss"])
			log.Println("Claims[IssuedAt]", claims["iat"])
			log.Println("Claims[ExpiresAt]", claims["exp"])
		} else {
			log.Println(err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
